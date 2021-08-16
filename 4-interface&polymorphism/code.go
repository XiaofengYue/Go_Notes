package main

import (
	"fmt"
	"strconv"
)

type Good interface {
	settleAccount() int
	orderInfo() string
}

type Phone struct {
	name     string
	quantity int
	price    int
}

type Gift struct {
	name     string
	quantity int
	price    int
}

func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}

func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

func (gift Gift) settleAccount() int {
	return 0
}

func (gift Gift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

func calculateAllPrice(goods []Good) int {
	all_price := 0
	for _, good := range goods {
		fmt.Printf(good.orderInfo())
		all_price += good.settleAccount()
	}
	return all_price
}

func main() {
	iphone := Phone{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}

	earpods := Gift{
		name:     "earnings",
		quantity: 1,
		price:    0,
	}

	goods := []Good{iphone, earpods}
	all_price := calculateAllPrice(goods)
	fmt.Printf("%d", all_price)
}
