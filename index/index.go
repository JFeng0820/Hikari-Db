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
