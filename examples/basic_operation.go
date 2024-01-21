package main

import (
	bitcask "bitcask-go"
	"fmt"
)

func main() {
	// 指定配置
	opts := bitcask.DefaultOptions
	opts.DirPath = "/tmp/bitcask-go"

	// 打开数据库
	db, err := bitcask.Open(opts)
	if err != nil {
		panic(err)
	}

	err = db.Put([]byte("name"), []byte("bitcask"))
	if err != nil {
		panic(err)
	}

	val, err := db.Get([]byte("name"))
	if err != nil {
		panic(err)
	}
	fmt.Println("val = ", string(val))

	err = db.Delete([]byte("name"))
	if err != nil {
		panic(err)
	}
}
