package main

import "fmt"

//上层业务逻辑层。去向下依赖抽象层
//中层抽象层。
//下层实现层，，去向上依赖抽象层

// ----> 抽象层
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// ----> 实现层
type BenZ struct {
}

func (b *BenZ) Run() {
	println("benz run")
}

type BenZDriver struct {
}

func (b *BenZDriver) Drive(car Car) {
	car.Run()
}

type ZhangSan struct {
}

func (z *ZhangSan) Drive(car Car) {
	fmt.Println("张三在驾驶 BenZ")
	car.Run()
}

// lisi
type LiSi struct {
}

func (l *LiSi) Drive(car Car) {
	fmt.Println("李四在驾驶 BenZ")
	car.Run()
}

// 王五
type WangWu struct {
}

func (w *WangWu) Drive(car Car) {
	fmt.Println("王五在驾驶 BenZ")
	car.Run()
}

// ----> 业务逻辑层
func main() {
	//zhanshankaiBenz
	var benz Car
	benz = new(BenZ)

	var zhang3D Driver
	zhang3D = new(ZhangSan)
	zhang3D.Drive(benz)

	//李四开宝马
	var li4D Driver
	li4D = new(LiSi)
	var bmw Car
	li4D.Drive(bmw)

	//王五开奔驰
	var wang5D Driver
	wang5D = new(WangWu)
	wang5D.Drive(benz)
}
