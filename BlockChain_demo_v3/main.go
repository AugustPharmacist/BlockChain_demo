package main

/*
1.定义结构
2.前区块哈希
3.当前区块哈希值
4.数据
5.创建区块
6.生成哈希
7.引入区块链
8.添加区块
9.重构代码
*/
func main() {
	/*	bc := NewBlockChain()
		bc.AddBlock("hello world !")
		bc.AddBlock("你好，中国。")

		for i, block := range bc.blocks {
			fmt.Printf("============== 区块高度：%d ===========\n", i)
			fmt.Printf("Version :%d\n", block.Version)
			fmt.Printf("PrevBlockHash :%x\n", block.PrevBlockHash)
			fmt.Printf("MerkelRoot :%x\n", block.MerkelRoot)
			timeFormat := time.Unix(int64(block.TimeStamp), 0).Format("15:04:05")
			fmt.Printf("TimeStamp :%s\n", timeFormat)
			//fmt.Printf("TimeStamp :%d\n",block.TimeStamp)
			fmt.Printf("Difficulty :%d\n", block.Difficulty)
			fmt.Printf("Nonce :%d\n", block.Nonce)
			fmt.Printf("Hash :%x\n", block.Hash)
			fmt.Printf("Data :%s\n", block.Data)
			pow := NewProofOfWork(*block)
			fmt.Printf("IsValid :%v\n", pow.IsVaild())

		}*/

	cli := CLI{}
	cli.Run()
}
