package main

import (
	"fmt"
	"time"
)

type Order struct {
	orderNo   int
	orderAmt  int
	orderDate time.Time
}
type Customer struct {
	custNumber  int
	custName    string
	phone       string
	creditLimit float32
	orders      []Order
}

func main() {

	c1 := Customer{101, "John Ley", "999982331", 2900.5, nil}
	c2 := Customer{102, "Whatever", "888999000", 345.5, nil}

	fmt.Println(c1)
	fmt.Println(c2)

	var c3 Customer = Customer{156, "Whatever", "888999000", 345.5, nil}
	fmt.Println(c3)

	var c4 Customer = Customer{custName: "Shivan LP"}
	fmt.Println(c4)

	c4.custNumber = 145
	c4.phone = "1000"
	c4.creditLimit = 9999
	fmt.Println(c4)

	newCust := Customer{
		custNumber: 133,
		custName:   "Rags",
		orders: []Order{
			Order{
				orderNo:   1,
				orderAmt:  100,
				orderDate: time.Now(),
			},
			Order{
				orderNo:   1,
				orderAmt:  100,
				orderDate: time.Now(),
			},
			Order{
				orderNo:   1,
				orderAmt:  100,
				orderDate: time.Now(),
			},
		},
	}
	fmt.Println(newCust.custNumber, newCust.custName)
	for _, order := range newCust.orders {
		fmt.Println(order.orderNo, order.orderAmt, order.orderDate.Day())
	}
}
