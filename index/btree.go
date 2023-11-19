package index

import (
	"bitcask-go/data"
	"bytes"
	"github.com/google/btree"
	"sort"
	"sync"
)

// Btree 索引，主要封装 Google 的BTree的库
// https://github.com/google/btree
type Btree struct {
	tree *btree.BTree
	lock *sync.RWMutex
}

// NewBTree 初始化 BTree 索引结构
func NewBTree() *Btree {
	return &Btree{
		tree: btree.New(32),
		lock: new(sync.RWMutex),
	}
}

func (bt *Btree) Put(key []byte, pos *data.LogRecordPos) bool {
	it := &Item{key: key, pos: pos}
	bt.lock.Lock()
	bt.tree.ReplaceOrInsert(it)
	bt.lock.Unlock()
	return true
}

func (bt *Btree) Get(key []byte) *data.LogRecordPos {
	it := &Item{key: key}
	btreeItem := bt.tree.Get(it)
	if btreeItem == nil {
		return nil
	}
	return btreeItem.(*Item).pos
}

func (bt *Btree) Delete(key []byte) bool {
	it := &Item{key: key}
	bt.lock.Lock()
	oldItem := bt.tree.Delete(it)
	if oldItem == nil {
		return false
	}
	bt.lock.Unlock()
	return true
}

func (bt *Btree) Size() int {
	return bt.tree.Len()
}

func (bt *Btree) Iterator(reverse bool) Iterator {
	if bt == nil {
		return nil
	}
	bt.lock.RLock()
	defer bt.lock.RUnlock()
	return newBTreeIterator(bt.tree, reverse)
}

func (bt *Btree) Close() error {
	return nil
}

type btreeIterator struct {
	currIndex int     // 当前索引下标
	reverse   bool    // 是否是反向遍历
	values    []*Item // key+位置索引信息
}

func newBTreeIterator(tree *btree.BTree, reverse bool) *btreeIterator {
	var idx int
	// 这里需要把数据都放到内存是因为我们的自定义的btree没有支持遍历的方法，若新索引有支持这不需要这样进行
	value := make([]*Item, tree.Len())

	// 将所有的数据存放到数组中
	saveValues := func(it btree.Item) bool {
		value[idx] = it.(*Item)
		idx++
		return true
	}
	if reverse {
		tree.Descend(saveValues)
	} else {
		tree.Ascend(saveValues)
	}

	return &btreeIterator{
		currIndex: 0,
		reverse:   reverse,
		values:    value,
	}
}

// Rewind 重新回到迭代器的起点
func (bit *btreeIterator) Rewind() {
	bit.currIndex = 0
}

// Seek 根据传入的key查找到第一个大于（或小于）等于的模板key，根据这个key开始遍历
func (bit *btreeIterator) Seek(key []byte) {
	if bit.reverse {
		// 二分查找法，标准库提供
		bit.currIndex = sort.Search(len(bit.values), func(i int) bool {
			return bytes.Compare(bit.values[i].key, key) <= 0
		})
	} else {
		bit.currIndex = sort.Search(len(bit.values), func(i int) bool {
			return bytes.Compare(bit.values[i].key, key) >= 0
		})
	}
}

// Next 跳转到下一个key
func (bit *btreeIterator) Next() {
	bit.currIndex += 1
}

// Valid 是否有效，即是否已经遍历完了所有的key，用于退出遍历
func (bit *btreeIterator) Valid() bool {
	return bit.currIndex < len(bit.values)
}

// Key 当前遍历位置的key数据
func (bit *btreeIterator) Key() []byte {
	return bit.values[bit.currIndex].key
}

// Value 当前遍历位置的Value数据
func (bit *btreeIterator) Value() *data.LogRecordPos {
	return bit.values[bit.currIndex].pos
}

// Close 关闭迭代器，释放相应资源
func (bit *btreeIterator) Close() {
	bit.values = nil
}
