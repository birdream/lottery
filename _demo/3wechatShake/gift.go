package main

var g1 = gift{
	id:       1,
	name:     "mobile",
	pic:      "",
	link:     "",
	gType:    giftTypeRealLarge,
	data:     "",
	dataList: nil,
	total:    2,
	left:     2,
	inUse:    true,
	rate:     1,
	rateMin:  0,
	rateMax:  0,
}

var g2 = gift{
	id:       2,
	name:     "charge",
	pic:      "",
	link:     "",
	gType:    giftTypeRealSmall,
	data:     "",
	dataList: nil,
	total:    5,
	left:     5,
	inUse:    true,
	rate:     10,
	rateMin:  0,
	rateMax:  0,
}

var g3 = gift{
	id:       3,
	name:     "Qp",
	pic:      "",
	link:     "",
	gType:    giftTypeCoupon,
	data:     "mall-c-2018",
	dataList: nil,
	total:    50,
	left:     50,
	inUse:    true,
	rate:     500,
	rateMin:  0,
	rateMax:  0,
}

var g4 = gift{
	id:       4,
	name:     "Qp50",
	pic:      "",
	link:     "",
	gType:    giftTypeCoupon,
	data:     "",
	dataList: []string{"c01", "c02", "c03", "c04", "c05", "c06", "c07", "c08", "c09", "c10", "c11"},
	total:    10,
	left:     10,
	inUse:    true,
	rate:     100,
	rateMin:  0,
	rateMax:  0,
}

var g5 = gift{
	id:       5,
	name:     "coin",
	pic:      "",
	link:     "",
	gType:    giftTypeCoin,
	data:     "10b",
	dataList: nil,
	total:    5,
	left:     5,
	inUse:    true,
	rate:     5000,
	rateMin:  0,
	rateMax:  0,
}
