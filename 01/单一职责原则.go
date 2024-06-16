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

func main() {
	c := &Clothes{}
	c.OnWork()
	c.OnShop()
}
