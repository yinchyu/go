package compent

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var (
	client = redis.NewClient(&redis.Options{
		Addr: "42.193.190.143:6388",
		DB:   1,
		//Password: "123456",
	})
)

func SetOp(number string) (string, error) {
	stringmap := client.HGetAll(number)
	usedtimes := stringmap.Val()["count"]
	code := stringmap.Val()["code"]
	times, _ := strconv.Atoi(usedtimes)
	if code != "" && times < 3 {
		client.HIncrBy(number, "count", 1)
		return code, nil
	} else if code != "" && times >= 3 {
		return "", fmt.Errorf("%s", "超过该时间段的获取次数")
	} else {
		codenumber := AuthCode(4)
		client.HSet(number, "code", codenumber)
		client.HSet(number, "count", 1)
		client.Expire(number, time.Minute*1)
		return codenumber, nil
	}
}
func SecKill() {
	router := gin.Default()
	router.GET("/:id", func(c *gin.Context) {
		// 返回一个json字符串
		value := c.Param("id")
		//获取到商品的id
		// 直接使用watch 会导致很多的碰撞
		e := client.Watch(func(tx *redis.Tx) error {
			remain, _ := client.Get(value).Int()
			if remain > 0 {
				client.Decr(value)
				c.JSON(200, gin.H{
					"code":  200,
					"id":    value,
					"state": true,
				})
			} else {
				c.JSON(200, gin.H{
					"code":  400,
					"id":    value,
					"state": false,
				})
			}
			return nil
		}, value)
		if e != nil {
			log.Println(e)
		}
	})

	router.Run(":80")
}

func AuthCode(n int) string {
	code := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		code = fmt.Sprintf("%s%d", code, rand.Intn(10))

	}

	return code

}

func AuthCode2() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%6v", rnd.Int31n(1000000))
	return vcode
}
