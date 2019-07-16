package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var mu sync.Mutex = sync.Mutex{}

type Prate struct {
	Rate  int // 万分之N中奖概率
	Total int // 总数量
	CodeA int // 中奖概率起始编码
	CodeB int // 中奖概率终止编码
	Left  *int32
}

var prizeList []string = []string{
	"1d",
	"2d",
	"3d",
	"",
}

var left = int32(1000)
var rateList []Prate = []Prate{
	Prate{Rate: 100, Total: 1000, CodeA: 0, CodeB: 9999, Left: &left},
	// Prate{Rate: 1, Total: 1, CodeA: 0, CodeB: 0, Left: &left},
	// Prate{Rate: 2, Total: 2, CodeA: 1, CodeB: 2, Left: &int32(2)},
	// Prate{Rate: 5, Total: 10, CodeA: 3, CodeB: 5, Left: &int32(10)},
	// Prate{Rate: 100, Total: 0, CodeA: 0, CodeB: 9999, Left: &int32(0)},
}

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func (c *lotteryController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return fmt.Sprintf("Big Wheel List: <br/> %s",
		strings.Join(prizeList, "<br/>\n"))
}

func (c *lotteryController) GetDebug() string {
	return fmt.Sprintf("Rate: %v \n", rateList)
}

func (c *lotteryController) GetPrize() string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	code := r.Intn(10000)

	var myPrize string
	var prizeRate *Prate

	for i, prize := range prizeList {
		rate := &rateList[i]
		if code >= rate.CodeA && code <= rate.CodeB {
			myPrize = prize
			prizeRate = rate
			break
		}
	}

	if myPrize == "" {
		myPrize = "Sorry, tyr it again!"

		return myPrize
	}

	if prizeRate.Total == 0 {
		return myPrize
	} else if *prizeRate.Left > 0 {
		// prizeRate.Left--
		left := atomic.AddInt32(prizeRate.Left, -1)
		if left >= 0 {
			return myPrize
		}
	}

	myPrize = "Sorry, tyr it again!!"
	return myPrize
}
