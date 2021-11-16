package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpcloud/tail"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	ok    bool
	wg    sync.WaitGroup
	tails *tail.Tail
)

func router() *gin.Engine {
	// 记录到文件。
	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	// 展示商品的基本信息
	r.GET("/:gid", show)
	r.GET("/handle", handle)
	r.GET("/handlewithlock", handlewithlock)
	r.GET("/handleone", handleone)
	r.GET("/handletwo", handletwo)
	r.GET("/handlechannel", handlechannel)
	r.GET("/log", logprocess)
	return r
}
func logprocess(c *gin.Context) {
	data, _ := os.ReadFile("./gin.log")
	c.DataFromReader()
	c.JSON(200, gin.H{
		"data": string(data),
	})

}
func show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("gid"))
	goods := mydb.FindGoodById(int64(id))
	covers := mydb.FindCoverByGoodsId(int64(id))
	details := mydb.FindDetailsByGoodsId(int64(id))
	params := mydb.FindParamsByGoodsId(int64(id))
	c.JSON(200, gin.H{
		"goods":   goods,
		"covers":  covers,
		"details": details,
		"params":  params,
	})

}

func handle(c *gin.Context) {
	gid := c.Query("gid")
	id, _ := strconv.Atoi(gid)
	seckillnum := 50
	wg.Add(seckillnum)
	InitializeSecKill(int64(id))
	for i := 0; i < seckillnum; i++ {
		userId := int64(i)
		go func() {
			err := HandleSeckill(int64(id), userId)
			if err != nil {
				fmt.Println("秒杀系统出错")
			} else {
				fmt.Printf("用户: %v抢购成功\n", userId)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	killedCount, err := GetKilledCount(int64(id))
	if err != nil {
		fmt.Println("秒杀系统出错")
	}
	fmt.Printf("一共秒杀出 %v 件商品", killedCount)
	c.JSON(200, gin.H{
		"state":   "failed",
		"message": "会存在超卖现象",
	})
}
func handlewithlock(c *gin.Context) {
	gid := c.Query("gid")
	id, _ := strconv.Atoi(gid)

	seckillNum := 56
	wg.Add(seckillNum)

	// 数据库中的商品、秒杀信息的初始化
	InitializeSecKill(int64(id))

	for i := 0; i < seckillNum; i++ {
		userId := int64(i)
		go func() {
			err := HandleSecKillWithLock(int64(id), userId)
			if err != nil {
				fmt.Println("秒杀系统出错")
			} else {
				fmt.Printf("用户: %v抢购成功\n", userId)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	killedCount, err := GetKilledCount(int64(id))
	if err != nil {
		fmt.Println("秒杀系统出错")
	}
	fmt.Printf("一共秒杀出 %v 件商品", killedCount)
	c.JSON(200, gin.H{
		"state":   "success",
		"message": "秒杀正常",
	})
}
func handleone(c *gin.Context) {
	gid := c.Query("gid")
	id, _ := strconv.Atoi(gid)

	seckillNum := 44
	wg.Add(seckillNum)

	// 数据库中的商品、秒杀信息的初始化
	InitializeSecKill(int64(id))

	for i := 0; i < seckillNum; i++ {
		userId := int64(i)
		go func() {
			err := HandleSecKillWithPccOne(int64(id), userId)
			if err != nil {
				fmt.Println("秒杀系统出错")
			} else {
				fmt.Printf("用户: %v抢购成功\n", userId)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	killedCount, err := GetKilledCount(int64(id))
	if err != nil {
		fmt.Println("秒杀系统出错")
	}
	fmt.Printf("一共秒杀出 %v 件商品", killedCount)
	c.JSON(200, gin.H{
		"state":   "failed",
		"message": "不知道什么原因，出现超卖现象",
	})
}
func handletwo(c *gin.Context) {
	gid := c.Query("gid")
	id, _ := strconv.Atoi(gid)
	seckillNum := 30
	wg.Add(seckillNum)
	// 数据库中的商品、秒杀信息的初始化
	InitializeSecKill(int64(id))

	for i := 0; i < seckillNum; i++ {
		userId := int64(i)
		go func() {
			err := HandleSecKillWithPccTwo(int64(id), userId)
			if err != nil {
				fmt.Println("秒杀系统出错")
			} else {
				fmt.Printf("用户: %v抢购成功\n", userId)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	killedCount, err := GetKilledCount(int64(id))
	if err != nil {
		fmt.Println("秒杀系统出错")
	}
	fmt.Printf("一共秒杀出 %v 件商品", killedCount)
	c.JSON(200, gin.H{
		"state":   "success",
		"message": "秒杀正常",
	})
}

func handlechannel(c *gin.Context) {
	gid := c.Query("gid")
	id, _ := strconv.Atoi(gid)
	seckillNum := 57
	// 数据库中的商品、秒杀信息的初始化
	InitializeSecKill(int64(id))
	go ChannelConsumer()
	for i := 0; i < seckillNum; i++ {
		userId := int64(i)
		go func() {
			err := HandleSecKillWithChannel(int64(id), userId)
			if err != nil {
				fmt.Println("秒杀系统出错")
			}
		}()
	}
	time.Sleep(time.Second * 10)
	killedCount, err := GetKilledCount(int64(id))
	if err != nil {
		fmt.Println("秒杀系统出错")
	}
	fmt.Printf("一共秒杀出 %v 件商品", killedCount)
	c.JSON(200, gin.H{
		"state":   "success",
		"message": "秒杀正常",
		"data":    "sdfsa",
	})
}
