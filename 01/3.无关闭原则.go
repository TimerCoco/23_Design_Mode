package main

import "fmt"

type Banker struct {
}

// Save 存款业务
func (b Banker) Save() {
	fmt.Println("存款业务")
}

// Loan 转账业务
func (b Banker) Loan() {
	fmt.Println("转账业务")
}

// Pay 支付业务
func (b Banker) Pay() {
	fmt.Println("支付业务")
}

func main() {
	banker := Banker{}
	banker.Save()
	banker.Loan()
	banker.Pay()
}
