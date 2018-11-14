package main

import "fmt"

/*
	1、定义结构
	2、前区块哈希
	3、当前区块哈希
	4、数据
	5、创建区块
	6、生成哈希
	7、引入区块链
	8、添加区块
	9、重构代码
*/
func main() {
	bc := NewBlockChain()
	bc.AddBlock("hello world")
	bc.AddBlock("你好，世界")

	for i, block := range bc.blocks {
		fmt.Printf("\n//////////////////////////////////////////////////////////////////////////////////////////\n\n")
		fmt.Printf("========= 区块链高度：%d =========\n", i)
		fmt.Printf("PrevBlockHash :%x\n", block.PrevBlockHash)
		fmt.Printf("Hash :%x\n", block.Hash)
		fmt.Printf("Data :%s\n", block.Data)
	}
}
