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

const (
	giftTypeCoin = iota
	giftTypeCoupon
	giftTypeCouponFix
	giftTypeRealSmall
	giftTypeRealLarge
)

type gift struct {
	id       int
	name     string
	pic      string
	link     string
	gType    int
	data     string
	dataList []string
	total    int // 0则不限制数量
	left     int
	inUse    bool
	rate     int // 万分之n 0-9999
	rateMin  int // 大于等于最小中奖编码
	rateMax  int // 小雨中奖编码
}

// 最大中奖号码
const rateMax = 10000

var mu sync.Mutex

var logger *log.Logger
var giftList []*gift

type lotteryController struct {
	Ctx iris.Context
}

func initLog() {
	f, err := os.Create("/var/log/lottery_demo.log")
	if err != nil {
		panic(err.Error())
	}
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func initGift() {
	giftList = make([]*gift, 5)

	giftList[0] = &g1
	giftList[1] = &g2
	giftList[2] = &g3
	giftList[3] = &g4
	giftList[4] = &g5

	rateStart := 0
	for _, data := range giftList {
		if !data.inUse {
			continue
		}

		data.rateMin = rateStart
		data.rateMax = rateStart + data.rate
		if data.rateMax >= rateMax {
			data.rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += data.rate
		}
	}
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	initLog()
	initGift()

	return app
}

func main() {
	mu = sync.Mutex{}

	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func (c *lotteryController) Get() string {
	count := 0
	total := 0
	for _, data := range giftList {
		if data.inUse && (data.total == 0 || data.total > 0 && data.left > 0) {
			count++
			total += data.left
		}
	}

	return fmt.Sprintf("Current having prize: %d, limit prize total: %d\n", count, total)
}

func (c *lotteryController) GetLucky() map[string]interface{} {
	mu.Lock()
	defer mu.Unlock()

	code := luckyCode()
	ok := false
	result := make(map[string]interface{})
	result["success"] = ok
	for _, data := range giftList {
		if !data.inUse || (data.total > 0 && data.left <= 0) {
			continue
		}

		if data.rateMin <= int(code) && data.rateMax > int(code) {
			sendData := ""

			switch data.gType {
			case giftTypeCoin:
				ok, sendData = sendCoin(data)
			case giftTypeCoupon:
				ok, sendData = sendCoupon(data)
			case giftTypeCouponFix:
				ok, sendData = sendCouponFix(data)
			case giftTypeRealLarge:
				ok, sendData = sendRealLarge(data)
			case giftTypeRealSmall:
				ok, sendData = sendRealSmall(data)
			}

			if ok {
				saveLuckyData(code, data.id, data.name, data.link, sendData, data.left)
				result["success"] = ok
				result["id"] = data.id
				result["name"] = data.name
				result["link"] = data.link
				result["data"] = sendData

				break
			}
		}
	}

	return result
}

func luckyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(int32(rateMax))

	return code
}
