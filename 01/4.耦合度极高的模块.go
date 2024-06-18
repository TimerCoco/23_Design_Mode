package main

import "fmt"

// 司机张三，李四，汽车，宝马，奔驰
// 1. 张三开奔驰
// 2. 李四开宝马
type Benz struct {
}

func (b Benz) run() {
	fmt.Println("benz run")
}

type BMW struct {
}

func (b BMW) run() {
	fmt.Println("bmw run")
}

type zhang3 struct {
}

func (z zhang3) drive(car *Benz) {
	car.run()
	fmt.Println("zhang3 drive benz")
}

// 李四开宝马
type li4 struct {
}

func (l li4) drive(car *BMW) {
	car.run()
	fmt.Println("li4 drive bmw")
}

func main() {
	benz := &Benz{}
	zhang3 := &zhang3{}
	zhang3.drive(benz)
	bmw := &BMW{}
	li4 := &li4{}
	li4.drive(bmw)
}
