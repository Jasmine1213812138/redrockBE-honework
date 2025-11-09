package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int //cangkushuliang
}

func (m Product) TotalValue() float64 {
	return m.Price * float64(m.Stock)
}
func (m Product) IsInStock() bool {
	return m.Stock > 0
}
func (m Product) Info() string {
	return fmt.Sprintf("商品: %s, 单价: %.2f, 库存: %d", m.Name, m.Price, m.Stock)
}
func (m *Product) Restock(amount int) {
	m.Stock += amount
}
func (m *Product) Sell(amount int) (success bool, message string) {
	if amount > m.Stock {
		success = false
		message = "失败，库存不足"
	} else {
		success = true
		m.Stock -= amount
		return success, fmt.Sprintf("message = “成功，剩余库存：%d", m.Stock)
	}
	return
}
func main() {
	book := Product{
		"Go编程书",
		89.5,
		10,
	}
	var amount int
	//第一次出售
	fmt.Println("欢迎光临，请问您想购买多少本书呢")
	_, err := fmt.Scan(&amount)
	for err != nil {
		fmt.Println("请确认一下您输入是否符合标准")
		_, err = fmt.Scan(&amount)
	}
	judge, message := book.Sell(amount)
	if judge {
		fmt.Println(message)
	} else {
		fmt.Println(message)
	}
	//进货20本
	book.Restock(20)
	//再次出售
	fmt.Println("欢迎光临，请问您想购买多少本书呢")
	_, err1 := fmt.Scan(&amount)
	for err1 != nil {
		fmt.Println("请确认一下您输入是否符合标准")
		_, err1 = fmt.Scan(&amount)
	}
	judge1, message1 := book.Sell(amount)
	if judge1 {
		fmt.Println(message1)
	} else {
		fmt.Println(message1)
	}
	//商品信息
	fmt.Printf("商品:%s,单价:%.2f,库存：%d\n", book.Name, book.Price, book.Stock)
	fmt.Printf("商品总价值：%.2f元", book.TotalValue())
}
