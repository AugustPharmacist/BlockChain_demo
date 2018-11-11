package main

import "fmt"

func main() {
	blockInterval := 21 //万
	reward := 50.0      //最初的奖励
	total := 0.0        //BTC总量

	//只要奖励大于0，就继续循环
	for reward > 0 {
		//4年内的挖矿总量
		amount := float64(blockInterval) * reward

		//所有的BTC累加
		total += amount

		//奖励衰减
		//reward /=2
		//为了优化程序，提高效率
		reward *= 0.5
	}
	fmt.Printf("reward: %f\n", reward)
	fmt.Printf("total: %f\n", total)
}
