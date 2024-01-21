# hikari-db介绍



## 什么是hikari-db

hikari-db是基于 Bitcask 存储模型，轻量、快速、可靠的 KV 存储引擎。

> 论文地址：<https://riak.com/assets/bitcask-intro.pdf>

bitcask 存储模型是由一个做分布式存储系统的商业化公司 Riak 提出来的，是一个简洁高效的存储引擎

## 设计概要

![image.png](docs/images/simple_design.jpg)

数据以追加的形式写入日志文件中。每个文件在达到预定大小后关闭，新的写入操作转移到新文件。

## 主要特点

### 优势

#### 写入优化

Bitcask 采用了追加写（append-only）的方式进行数据存储。所有的写操作都直接追加到文件的末尾，极大地提高了写入速度，并减少了磁盘寻址时间。

#### 快速查找

Bitcask 通过在内存中维护一个哈希表来实现快速的键查找。每个键都指向其对应值的磁盘位置，从而使得读取操作非常快速。

#### 高效的空间利用

Bitcask 定期执行合并操作（compaction），在此过程中，它会删除过期或重复的数据，从而优化存储空间的使用。



### 缺点

#### key 必须在内存中维护

所有 key 保留在内存中，这意味着您的系统必须具有足够的内存来容纳所有的 key。



## 快速上手

### 基本操作

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





## 现状

本项目还处于待完善状态，仅可作为学习项目使用，欢迎大家进行交流



