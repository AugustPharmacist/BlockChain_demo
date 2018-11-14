package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {

	//1.打开数据库（如果没有会创建）
	//110
	//rwx
	db, err := bolt.Open("src/BlockChain_demo/BlockChain_demo_v3/demo/test.db", 0600, nil)

	if err != nil {
		log.Panic(err)
	}

	//2.关闭数据库
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		//3.找到我们的桶，通过桶的名字
		//returns nil if the bucket does not exist.
		bucket := tx.Bucket([]byte("bucketName1"))

		//如果没有找到，先创建
		bucket, err := tx.CreateBucketIfNotExists([]byte("bucketName1"))
		if err != nil {
			log.Panic(err)
		}

		//4.写数据
		err = bucket.Put([]byte("1"), []byte("hello"))
		if err != nil {
			log.Panic(err)
		}

		err = bucket.Put([]byte("2"), []byte("world"))
		if err != nil {
			log.Panic(err)
		}

		//5.读数据
		data1 := bucket.Get([]byte("1"))
		data2 := bucket.Get([]byte("2"))
		data3 := bucket.Get([]byte("3"))

		fmt.Printf("data 1 : %s\n", data1)
		fmt.Printf("data 2 : %s\n", data2)
		fmt.Printf("data 3 : %s\n", data3)

		return nil
	})

}
