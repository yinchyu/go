package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type MyDB struct {
	DB *sqlx.DB
}

func InitDB() (*MyDB, error) {
	dsn := "root:123456@tcp(www.yinchangyu.top:3306)/seckill?charset=utf8mb4&parseTime=true"
	DB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed err:%v\n", err)
	}
	return &MyDB{
		DB: DB,
	}, DB.Ping()
}

func (mydb *MyDB) CloseDB() {
	err := mydb.DB.Close()
	if err != nil {
		panic(err)
	}
}

func (mydb *MyDB) FindGoodById(GoodId int64) *Goods {
	sqlstr := "select * from t_goods where goods_id = ?"
	newgoods := &Goods{}
	err := mydb.DB.Get(newgoods, sqlstr, GoodId)
	if err != nil {
		log.Println(err)
	}
	return newgoods
}

func (mydb *MyDB) FindCoverByGoodsId(GoodId int64) []*GoodsCover {
	sqlstr := "select * from t_goods_cover where goods_id=? order by gc_order"
	var goodsCover []*GoodsCover
	err := mydb.DB.Select(&goodsCover, sqlstr, GoodId)
	if err != nil {
		log.Println(err)
	}
	return goodsCover
}

func (mydb *MyDB) FindDetailsByGoodsId(goodsId int64) []GoodsDetail {
	sqlStr := "select * from t_goods_detail where goods_id = ? order by gd_order"
	var goodsDetails []GoodsDetail
	err := mydb.DB.Select(&goodsDetails, sqlStr, goodsId)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
	}
	return goodsDetails
}
func (mydb *MyDB) FindParamsByGoodsId(goodsId int64) []*GoodsParam {
	sqlStr := "select * from t_goods_param where goods_id = ? order by gp_order"
	var goodsParams []*GoodsParam
	err := mydb.DB.Select(&goodsParams, sqlStr, goodsId)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
	}
	return goodsParams
}

func (mydb *MyDB) UpdateCountByGoodsId(goodsId int64) error {
	sqlStr := "update t_promotion_seckill set ps_count = 20, version = 0 where goods_id = ?"
	_, err := mydb.DB.Exec(sqlStr, goodsId)
	if err != nil {
		return err
	}
	return nil
}

// 获取商品的数量
func (mydb *MyDB) SelectCountByGoodsId(goodsId int64) (int64, error) {
	sqlStr := "select ps_count from t_promotion_seckill where goods_id = ?"
	var count int64
	err := mydb.DB.Get(&count, sqlStr, goodsId)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (mydb *MyDB) SelectCountByGoodsIdPcc(goodsId int64) (int64, error) {
	// 增加了for update 保证数据的获取是互斥的操作，但是for update 之后数据会自动进行释放，所以并不能起到保护的作用。
	// 正确的做法是在获取数据的时候就进行加锁控制
	sqlStr := "select ps_count from t_promotion_seckill where goods_id = ? for update"
	var count int64
	err := mydb.DB.Get(&count, sqlStr, goodsId)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (mydb *MyDB) ReduceStockByGoodsId(goodsId int64, count int64) error {
	sqlStr := "update t_promotion_seckill set ps_count = ? where goods_id = ?"
	_, err := mydb.DB.Exec(sqlStr, count, goodsId)
	if err != nil {
		return err
	}
	return nil
}

func (mydb *MyDB) ReduceByGoodsId(goodsId int64) (int64, error) {
	sqlStr := "update t_promotion_seckill set ps_count = ps_count-1 where ps_count>0 and goods_id = ?"
	res, err := mydb.DB.Exec(sqlStr, goodsId)
	count, getErr := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if getErr != nil {
		return 0, getErr
	}
	return count, nil
}

func (mydb *MyDB) SelectGoodByGoodsId(goodsId int64) (PromotionSecKill, error) {
	sqlStr := "select * from t_promotion_seckill where goods_id = ?"
	var promotion PromotionSecKill
	err := mydb.DB.Get(&promotion, sqlStr, goodsId)
	if err != nil {
		return promotion, err
	}
	return promotion, nil
}

func (mydb *MyDB) ReduceStockByOcc(goodsId int64, num int64, version int64) (int64, error) {
	sqlStr := "update t_promotion_seckill set ps_count = ps_count-?, version = version+1 " +
		"where version = ? and goods_id = ?"
	res, err := mydb.DB.Exec(sqlStr, num, version, goodsId)
	count, getErr := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if getErr != nil {
		return 0, getErr
	}
	return count, nil
}
func (mydb *MyDB) DeleteByGoodsId(goodsId int64) error {
	sqlstr := "delete from t_success_killed where goods_id=?"
	_, err := mydb.DB.Exec(sqlstr, goodsId)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func (mydb *MyDB) CreatOrder(killed SuccessKilled) error {
	sqlStr := "insert into t_success_killed (goods_id, user_id, state, create_time) values (?, ?, ?, ?)"
	_, err := mydb.DB.Exec(sqlStr, killed.GoodsId, killed.UserId, killed.State, killed.CreateTime)
	if err != nil {
		return err
	}
	return nil
}

func (mydb *MyDB) GetKilledCountByGoodsId(goodsId int64) (int64, error) {
	sqlStr := "select count(*) from t_success_killed where goods_id = ?"
	var count int64
	err := mydb.DB.Get(&count, sqlStr, goodsId)
	if err != nil {
		return count, err
	}
	return count, nil
}
