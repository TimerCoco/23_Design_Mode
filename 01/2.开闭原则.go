package main

import "fmt"

// AbstractBanker 抽象的银行业务员
type AbstractBanker interface {
	DbBusi() //抽象的处理业务接口
}

// SaveBanker 存款的业务员
type SaveBanker struct {
	//相当于继承AbstractBanker
}

func (b *SaveBanker) DbBusi() {
	fmt.Println("存款业务")
}

//**********************************

// TransferBanker 转账的业务员
type TransferBanker struct {
	//相当于继承AbstractBanker
}

func (b *TransferBanker) DbBusi() {
	fmt.Println("转账业务")
}

// BankBusiness Bank 设计一个架构层，专门负责处理业务
func BankBusiness(banker AbstractBanker) {
	//通过接口向下调用从而实现多态
	banker.DbBusi()
}
func main() {
	////存款业务员
	//banker := SaveBanker{}
	//banker.DbBusi()
	////转账业务员
	//banker2 := TransferBanker{}
	//banker2.DbBusi()

	//存款业务
	BankBusiness(&SaveBanker{})
	//转账业务
	BankBusiness(&TransferBanker{})

}

//func main() {
//	var banker AbstractBanker
//	banker = new(SaveBanker)
//	banker.DbBusi()
//	banker = new(TransferBanker)
//	banker.DbBusi()
//}
