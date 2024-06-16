package main

import "fmt"

type Clothes struct {
	//穿衣服
}

// 工作穿衣
func (c *Clothes) OnWork() {
	//工作穿衣服
	fmt.Println("工作穿衣服")
}
func (c *Clothes) OnShop() {
	fmt.Println("购物穿衣服")
}

// 这是一个错误的案例，因为这个类里面有2个方法，如果以后有新需求，就要修改这个类，
func main() {
	c := &Clothes{}
	c.OnWork()
	c.OnShop()
}
