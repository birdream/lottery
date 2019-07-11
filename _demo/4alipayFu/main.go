package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type gift struct {
	id      int
	name    string
	pic     string
	link    string
	inUse   bool
	rate    int // 万分之n 0-9999
	rateMin int // 大于等于最小中奖编码
	rateMax int // 小雨中奖编码
}

// 最大中奖号码
const rateMax = 11

var mu sync.Mutex

var logger *log.Logger
var giftList []*gift

type lotteryController struct {
	Ctx iris.Context
}

func initLog() {
	f, err := os.Create("/var/log/alipayFu_demo.log")
	if err != nil {
		panic(err.Error())
	}
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	return app
}

func main() {
	initLog()

	mu = sync.Mutex{}

	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func (c *lotteryController) Get() string {
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	giftList := giftRage(rate)

	return fmt.Sprintf("%v\n", giftList)
}

func luckyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(int32(rateMax))
	fmt.Print(code)

	return code
}

func (c *lotteryController) GetLucky() map[string]interface{} {
	rate := c.Ctx.URLParamDefault("rate", "4,3,2,1,0")
	uid, _ := c.Ctx.URLParamInt("uid")

	code := luckyCode()
	ok := false
	result := make(map[string]interface{})
	result["success"] = ok
	giftList := giftRage(rate)

	for _, data := range giftList {
		if !data.inUse {
			continue
		}

		if data.rateMin <= int(code) && data.rateMax > int(code) {
			sendData := data.pic

			saveLuckyData(code, data.id, data.name, data.link, sendData)
			result["success"] = ok
			result["id"] = data.id
			result["name"] = data.name
			result["link"] = data.link
			result["data"] = sendData
			result["uid"] = uid

			break
		}
	}

	return result
}
