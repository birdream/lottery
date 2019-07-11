/**
 * curl http://localhost:8080/set\?uid\=1\&money\=100\&num\=100
 * curl http://localhost:8080
 * wrk -t10 -c10 -d5 http://localhost:8080/get\?uid\=3\&id\=1294913411
 */
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type lottertController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lottertController{})

	return app
}

func main() {
	go fetchPLMoney()

	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func (c *lottertController) Get() map[uint32][2]int {
	rs := make(map[uint32][2]int)
	// for id, list := range packageList {
	// 	var money int
	// 	for _, v := range list {
	// 		money += int(v)
	// 	}

	// 	rs[id] = [2]int{len(list), money}
	// }

	packageList.Range(func(key, value interface{}) bool {
		id := key.(uint32)
		list := value.([]uint)

		var money int
		for _, v := range list {
			money += int(v)
		}

		rs[id] = [2]int{len(list), money}
		return true
	})

	return rs
}

// /set?uid=1&money=100&num=10
func (c *lottertController) GetSet() string {
	uid, errUID := c.Ctx.URLParamInt("uid")
	money, errMoney := c.Ctx.URLParamFloat64("money")
	num, errNum := c.Ctx.URLParamInt("num")

	if errMoney != nil || errMoney != nil || errNum != nil {
		return fmt.Sprintf("type error: %d %d %d \n", errUID, errNum, errMoney)
	}

	moneyTotal := int(money * 100)
	if uid < 1 || moneyTotal < num || num < 1 {
		return fmt.Sprintf("value error: %d %d %d \n", errUID, errNum, errMoney)
	}

	// 金额分配算法
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rMax := 0.55 // 随机最大值
	if num > 1000 {
		rMax = 0.01
	} else if num >= 100 {
		rMax = 0.1
	} else if num >= 10 {
		rMax = 0.3
	}
	list := make([]uint, num)

	leftMoney := moneyTotal
	leftNum := num
	//大循环开始分配金额到每个红包
	for leftNum > 0 {
		if leftNum == 1 {
			//剩余都给他
			list[num-1] = uint(leftMoney)
			break
		}

		if leftMoney == leftNum {
			for i := num - leftNum; i < num; i++ {
				list[i] = 1
			}
			break
		}

		rMoney := int(float64(leftMoney-leftNum) * rMax)
		m := r.Intn(rMoney)
		if m < 1 {
			m = 1
		}
		list[num-leftNum] = uint(m)
		leftMoney -= m
		leftNum--
	}

	id := r.Uint32()
	// packageList[id] = list
	packageList.Store(id, list)

	return fmt.Sprintf("/get?id=%d&uid=%d&num=%d", id, uid, num)
}

// /get?id=1&uid=1
func (c *lottertController) GetGet() string {
	uid, errUID := c.Ctx.URLParamInt("uid")
	id, errID := c.Ctx.URLParamInt("id")
	if errID != nil || errUID != nil {
		return fmt.Sprintf("err %d, %d", errUID, errID)
	}

	if uid < 1 || id < 1 {
		return fmt.Sprintf("err %d, %d", uid, id)
	}

	// list, ok := packageList[uint32(id)]
	list1, ok := packageList.Load(uint32(id))
	list := list1.([]uint)

	if !ok || len(list) < 1 {
		return fmt.Sprintf("Red Pod Not Exist, id: %d", id)
	}

	//分配一个随机数
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// i := r.Intn(len(list))
	// money := list[i]

	// if len(list) > 1 {
	// 	if i == len(list)-1 {
	// 		// packageList[uint32(id)] = list[:i]
	// 		packageList.Store(uint32(id), list[:i])
	// 	} else if i == 0 {
	// 		// packageList[uint32(id)] = list[1:i]
	// 		packageList.Store(uint32(id), list[1:i])
	// 	} else {
	// 		// packageList[uint32(id)] = append(list[:i], list[i+1:]...)
	// 		packageList.Store(uint32(id), append(list[:i], list[i+1:]...))
	// 	}
	// } else {
	// 	// delete(packageList, uint32(id))
	// 	packageList.Delete(uint32(id))
	// }

	// 构造一个抢红包任务
	callback := make(chan uint)
	t := task{id: uint32(id), callback: callback}
	chTasks <- t
	money := <-callback
	if money <= 0 {
		return "Sorry, no red pod"
	} else {
		return fmt.Sprintf("Congres!, money: %d", money)
	}

}
