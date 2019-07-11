package main

import (
	"fmt"
	"strconv"
	"strings"
)

var g1 = gift{
	id:      1,
	name:    "富强福",
	pic:     "富强福.jpg",
	link:    "",
	inUse:   true,
	rate:    0,
	rateMin: 0,
	rateMax: 0,
}

var g2 = gift{
	id:      2,
	name:    "和谐福",
	pic:     "和谐福.jpg",
	link:    "",
	inUse:   true,
	rate:    0,
	rateMin: 0,
	rateMax: 0,
}

var g3 = gift{
	id:      3,
	name:    "友善福",
	pic:     "友善福.jpg",
	link:    "",
	inUse:   true,
	rate:    0,
	rateMin: 0,
	rateMax: 0,
}

var g4 = gift{
	id:      4,
	name:    "爱国福",
	pic:     "爱国福.jpg",
	link:    "",
	inUse:   true,
	rate:    0,
	rateMin: 0,
	rateMax: 0,
}

var g5 = gift{
	id:      5,
	name:    "敬业福",
	pic:     "敬业福.jpg",
	link:    "",
	inUse:   true,
	rate:    0,
	rateMin: 0,
	rateMax: 0,
}

func newGift() *[5]gift {
	giftList := new([5]gift)

	giftList[0] = g1
	giftList[1] = g2
	giftList[2] = g3
	giftList[3] = g4
	giftList[4] = g5

	return giftList
}

func giftRage(rate string) *[5]gift {
	giftList := newGift()
	rates := strings.Split(rate, ",")
	ratesLen := len(rates)
	rateStart := 0

	for i, data := range giftList {
		if !data.inUse {
			continue
		}

		gRate := 0
		var err error

		if i < ratesLen {
			gRate, err = strconv.Atoi(rates[i])
			check(err)
		}
		giftList[i].rate = gRate
		giftList[i].rateMin = rateStart
		giftList[i].rateMax = rateStart + gRate
		if giftList[i].rateMax >= rateMax {
			giftList[i].rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += gRate
		}
	}

	fmt.Printf("giftList: %v \n", giftList)

	return giftList
}

func saveLuckyData(code int32, id int, name string, link string, sendData string) {
	logger.Printf("lucky, code: %d, gift: %d, name: %s, link: %s, data: %s",
		code, id, name, link, sendData)
}
