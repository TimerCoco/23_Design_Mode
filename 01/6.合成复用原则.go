package main

import "fmt"

//能用合成就不要用继承

type cat struct {
}

func (c cat) eat() {
	fmt.Println("cat eat")
}

// 添加一个睡觉的方法：用继承实现
type catSleep struct {
	cat
}

func (c catSleep) Sleep() {
	fmt.Println("cat sleep")
}

// 添加一个睡觉的方法：用组合实现
type catSleep2 struct {
	C *cat //组合一个cat类
}

func (c catSleep2) eat() {
	c.C.eat() //调用组合进来的
}
func (c catSleep2) sleep() {
	fmt.Println("小猫睡觉")
}

func main() {
	c := cat{}
	c.eat()
	fmt.Println("--------------------------")
	cSleep := catSleep{c}
	cSleep.Sleep()
	cSleep.eat()

	fmt.Println("--------------------------")
	// 组合实现
	cc := &catSleep2{}
	cc.eat()
	cc.sleep()

}
