package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	lock     sync.Mutex
	instance *singleton
	once     sync.Once
)

type singleton chan [2]int64

func InitializeSecKill(goodsId int64) {
	tx, err := mydb.DB.Beginx()
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
	}
	err1 := mydb.DeleteByGoodsId(goodsId)
	err2 := mydb.UpdateCountByGoodsId(goodsId)
	if err1 != nil {
		fmt.Println(err1)
		tx.Rollback()
	}
	if err2 != nil {
		fmt.Println(err2)
		tx.Rollback()
	}
	tx.Commit()
}

func HandleSeckill(goodsId int64, userId int64) error {
	tx, err := mydb.DB.Beginx()
	if err != nil {
		return err
	}

	// 检查库存
	count, errCount := mydb.SelectCountByGoodsId(goodsId)
	if errCount != nil {
		return errCount
	}

	if count > 0 {
		// 1.扣库存
		errRed := mydb.ReduceStockByGoodsId(goodsId, count-1)
		if errRed != nil {
			tx.Rollback()
			return errRed
		}

		// 2.创建订单
		killed := SuccessKilled{
			GoodsId:    goodsId,
			UserId:     userId,
			State:      0,
			CreateTime: time.Now(),
		}
		errCre := mydb.CreatOrder(killed)
		if errCre != nil {
			tx.Rollback()
			return errCre
		}
	}
	tx.Commit()
	return nil
}

func HandleSecKillWithLock(goodsId int64, userId int64) error {
	// 完全的降低了并发量, 因为只有一把锁，所有的进程都会抢占这一把锁，所以会导致资源冲突
	lock.Lock()
	err := HandleSeckill(goodsId, userId)
	lock.Unlock()
	return err
}

func HandleSecKillWithPccOne(goodsId int64, userId int64) error {
	tx, err := mydb.DB.Beginx()
	if err != nil {
		return err
	}

	// 检查库存
	count, errCount := mydb.SelectCountByGoodsIdPcc(goodsId)
	if errCount != nil {
		tx.Rollback()
		return errCount
	}

	if count > 0 {
		// 1.扣库存
		errRed := mydb.ReduceStockByGoodsId(goodsId, count-1)
		if errRed != nil {
			tx.Rollback()
			return errRed
		}

		// 2.创建订单
		killed := SuccessKilled{
			GoodsId:    goodsId,
			UserId:     userId,
			State:      0,
			CreateTime: time.Now(),
		}
		errCre := mydb.CreatOrder(killed)
		if errCre != nil {
			tx.Rollback()
			return errCre
		}
	}
	tx.Commit()
	return nil
}

func HandleSecKillWithPccTwo(goodsId int64, userId int64) error {
	tx, err := mydb.DB.Beginx()
	if err != nil {
		return err
	}

	// 1.扣库存，加锁
	count, errCount := mydb.ReduceByGoodsId(goodsId)
	if errCount != nil {
		tx.Rollback()
		return errCount
	}

	if count > 0 {
		// 2.创建订单
		killed := SuccessKilled{
			GoodsId:    goodsId,
			UserId:     userId,
			State:      0,
			CreateTime: time.Now(),
		}
		errCre := mydb.CreatOrder(killed)
		if errCre != nil {
			tx.Rollback()
			return errCre
		}
	}
	tx.Commit()
	return nil
}

func HandleSecKillWithOcc(goodsId int64, userId int64, num int64) error {
	tx, err := mydb.DB.Beginx()
	if err != nil {
		return err
	}

	// 检查库存
	promotion, errCount := mydb.SelectGoodByGoodsId(goodsId)
	if errCount != nil {
		tx.Rollback()
		return errCount
	}

	if promotion.PsCount >= num {
		// 1.扣库存
		count, errRed := mydb.ReduceStockByOcc(goodsId, num, promotion.Version)
		if errRed != nil {
			tx.Rollback()
			return errRed
		}

		if count > 0 {
			// 2.创建订单
			killed := SuccessKilled{
				GoodsId:    goodsId,
				UserId:     userId,
				State:      0,
				CreateTime: time.Now(),
			}
			errCre := mydb.CreatOrder(killed)
			if errCre != nil {
				tx.Rollback()
				return errCre
			}
		} else {
			tx.Rollback()
		}
	} else {
		tx.Rollback()
		return errors.New("库存不够了")
	}
	tx.Commit()
	return nil
}

func HandleSecKillWithChannel(goodsId int64, userId int64) error {
	kill := [2]int64{goodsId, userId}
	chann := GetInstance()
	*chann <- kill
	return nil
}

func ChannelConsumer() {
	for {
		kill, ok := <-(*GetInstance())
		if !ok {
			continue
		}
		err := HandleSeckill(kill[0], kill[1])
		if err != nil {
			fmt.Println("秒杀出错")
		} else {
			fmt.Printf("用户: %v 秒杀成功\n", kill[1])
		}
	}
}

func GetKilledCount(goodsId int64) (int64, error) {
	return mydb.GetKilledCountByGoodsId(goodsId)
}

func GetInstance() *singleton {
	once.Do(func() {
		// 容量视库存大小而定，这里设为100
		ret := make(singleton, 100)
		instance = &ret
	})
	return instance
}
