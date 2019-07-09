/*
 * 1 即开即得
 * 2 双色球自选
 */
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

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

// 即开即得
func (c *lotteryController) Get() string {
	var prize string

	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Intn(10)

	switch {
	case code == 1:
		prize = "一等"
	case code >= 2 && code <= 3:
		prize = "二等"
	case code >= 4 && code <= 6:
		prize = "三等"
	default:
		return fmt.Sprintf("code=%d <br/> Sorry", code)
	}
	return fmt.Sprintf("code=%d <br/>Congress! your prize is %s", code, prize)
}

// 双色球自选
func (c *lotteryController) GetPrize() string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	var prize [7]int

	// 6个红色球 1-33
	for i := 0; i < 6; i++ {
		prize[i] = r.Intn(33) + 1
	}

	// 最后一位的蓝色球 1-16
	prize[6] = r.Intn(16) + 1

	return fmt.Sprintf("Prize in today is: %v", prize)
}
