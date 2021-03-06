package main

import (
	"fmt"
	"os"
)

const Usage = `
	./blockchain createBlockChain "创建区块链"
	./blockchain addBlock DATA   "添加数据到区块链"
	./blockchain printChain "打印区块链"
`

//定义一个CLI，里面包含BlockChain，所有细节工作交给bc，命令的解析工作交给CLI
type CLI struct {
	//bc *BlockChain
}

//定义一个run函数，负责接收命令行的数据，然后根据命令进行解析，并完成最终的调用
func (cli *CLI) Run() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("输入参数个数错误，请检查！\n")
		fmt.Println(Usage)
		os.Exit(0)
	}
	cmd := args[1]

	switch cmd {
	case "createBlockChain":
		fmt.Println("创建区块链命令被调用！\n")
		bc := CreateBlockChain()
		defer bc.Db.Close()

	case "addBlock":
		fmt.Println("添加区块连命令被调用！\n")
		bc := NewBlockChain()
		defer bc.Db.Close()

		//1.检查参数个数
		if len(args) == 3 {
			//2.获取数据
			data := args[2]
			//3.调用真正的添加区块函数
			bc.AddBlock(data)
		} else {
			fmt.Println("参数无效！")
			fmt.Println(Usage)
		}

	case "printChain":
		fmt.Println("打印区块命令被调用\n")
		bc := NewBlockChain()
		defer bc.Db.Close()
		bc.PrintChain()

	default:
		fmt.Println("无效的命令，请检查！\n")
	}
}
