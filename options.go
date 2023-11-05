package bitcask

import "os"

type Options struct {
	// 数据库文件目录
	DirPath string

	// 数据文件的大小
	DataFileSize int64

	// 每次写数据是否持久化
	SyncWrites bool

	// 索引类型
	IndexType IndexerType
}

// IteratorOption 索引迭代器配置项
type IteratorOption struct {
	// 遍历数据为指定值的 Key，默认为空
	Prefix []byte
	// 是否反向遍历，默认为 false 是正向
	Reverse bool
}

type IndexerType = int8

const (
	// BTree 索引
	BTree IndexerType = iota + 1

	// ART Adpative Radix Tree 自适应基数树索引
	ART
)

var DefaultOptions = Options{
	DirPath:      os.TempDir(),
	DataFileSize: 256 * 1024 * 1024, // 256MB
	SyncWrites:   false,
	IndexType:    BTree,
}

var DefaultIteratorOptions = IteratorOption{
	Prefix:  nil,
	Reverse: false,
}
