/*
 * curl http://localhost:8080
 * curl --data "users=norman,norman2,jay,mingy" http://localhost:8080/import
 * curl http://localhost:8080/lucky
 *
 */
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var userList []string
var mu sync.Mutex

type lottryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lottryController{})
	return app
}

func main() {
	app := newApp()
	// userList = make([]string, 0)
	userList = []string{}
	mu = sync.Mutex{}

	app.Run(iris.Addr(":8080"))
}

func (c *lottryController) Get() string {
	count := len(userList)
	return fmt.Sprintf("当前参与用户： %d\n", count)
}

func (c *lottryController) PostImport() string {
	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",")

	mu.Lock()
	defer mu.Unlock()

	count1 := len(userList)

	for _, u := range users {
		u = strings.TrimSpace(u)
		if len(u) > 0 {
			userList = append(userList, u)
		}
	}
	count2 := len(userList)

	return fmt.Sprintf("当前参与人数：%d， 成功倒入人数：%d \n", count2, (count2 - count1))
}

func (c *lottryController) GetLucky() string {
	mu.Lock()
	defer mu.Unlock()

	count := len(userList)
	if count > 1 {
		seed := time.Now().UnixNano()
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count))

		user := userList[index]
		userList = append(userList[0:index], userList[index+1:]...)

		return fmt.Sprintf("当前中奖用户：%s， 剩余用户数：%d \n", user, count-1)
	} else if count == 1 {
		user := userList[0]

		return fmt.Sprintf("当前中奖用户：%s， 剩余用户数：%d \n", user, count-1)
	} else {
		return fmt.Sprintf("没有用户啦 需要倒入一些")
	}
}
