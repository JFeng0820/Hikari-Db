package main

import (
	bitcask "bitcask-go"
	bitcask_redis "bitcask-go/redis"
	"github.com/tidwall/redcon"
	"log"
	"sync"
)

// TODO 用配置项
const addr = "127.0.0.1:6380"

type BitcaskServer struct {
	dbs    map[int]*bitcask_redis.RedisDataStructure
	server *redcon.Server
	mu     sync.RWMutex
}

func main() {
	// 打开 Redis 数据结构和服务
	redisDataStructure, err := bitcask_redis.NewRedisDataStructure(bitcask.DefaultOptions)
	if err != nil {
		panic(err)
	}

	// 初始化 BitcaskServer
	bitcaskServer := &BitcaskServer{
		dbs: make(map[int]*bitcask_redis.RedisDataStructure),
	}
	bitcaskServer.dbs[0] = redisDataStructure

	// 初始化redis服务端
	bitcaskServer.server = redcon.NewServer(addr, execClientCommand, bitcaskServer.accept, bitcaskServer.close)
	bitcaskServer.listen()
}

func (svr *BitcaskServer) listen() {
	log.Printf("bitcask server running, ready to accept connections.")
	_ = svr.server.ListenAndServe()
}

func (svr *BitcaskServer) close(conn redcon.Conn, err error) {
	for _, db := range svr.dbs {
		_ = db.Close()
	}
	_ = svr.server.Close()
}

func (svr *BitcaskServer) accept(conn redcon.Conn) bool {
	cli := new(BitcaskClient)
	svr.mu.Lock()
	defer svr.mu.Unlock()
	cli.server = svr
	// TODO 之后接收用户的配置项
	cli.db = svr.dbs[0]
	conn.SetContext(cli)
	return true
}

//func main() {
//	conn, err := net.Dial("tcp", "localhost:6379")
//	if err != nil {
//		panic(err)
//	}
//	defer conn.Close()
//
//	// 发送命令
//	cmd := "set key hello \r\n"
//	conn.Write([]byte(cmd))
//
//	// 返回响应
//	reader := bufio.NewReader(conn)
//	res, err := reader.ReadString('\n')
//	fmt.Println(res, err)
//}
