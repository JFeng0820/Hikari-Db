package index

import (
	"bitcask-go/data"
	"bytes"
	"github.com/google/btree"
)

type Indexer interface {
	// Put 向索引中存储 key 对应的数据位置信息
	Put(key []byte, pos *data.LogRecordPos) bool

	// Get 根据 key 取出的索引位置信息
	Get(key []byte) *data.LogRecordPos

	// Delete 根据 key 删除的索引位置信息
	Delete(key []byte) bool

	// Size 索引中的数据量
	Size() int

	// Iterator 索引迭代器
	Iterator(reverse bool) Iterator
}

type IndexType = int8

const (
	// Btree 索引
	BTree IndexType = iota + 1

	// ART 自适应基数树索引
	ART
)

func NewIndexer(typ IndexType) Indexer {
	switch typ {
	case BTree:
		return NewBTree()
	case ART:
		return nil
	default:
		panic("unsupported index tyoe")
	}

}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}

type Iterator interface {
	// Rewind 重新回到迭代器的起点
	Rewind()

	// Seek 根据传入的key查找到第一个大于（或小于）等于的模板key，根据这个key开始遍历
	Seek(key []byte)

	// Next 跳转到下一个key
	Next()

	// Valid 是否有效，即是否已经遍历完了所有的key，用于退出遍历
	Valid() bool

	// Key 当前遍历位置的key数据
	Key() []byte

	// Value 当前遍历位置的Value数据
	Value() *data.LogRecordPos

	// Close 关闭迭代器，释放相应资源
	Close()
}
