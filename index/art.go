package index

import (
	"bitcask-go/data"
	"bytes"
	goart "github.com/plar/go-adaptive-radix-tree"
	"sort"
	"sync"
)

// AdaptiveRadixTree 自适应基数树索引
// 主要封装了 https://github.com/plar/go-adaptive-radix-tree
type AdaptiveRadixTree struct {
	tree goart.Tree
	lock *sync.RWMutex
}

func NewART() *AdaptiveRadixTree {
	return &AdaptiveRadixTree{
		tree: goart.New(),
		lock: new(sync.RWMutex),
	}
}

func (art *AdaptiveRadixTree) Put(key []byte, pos *data.LogRecordPos) *data.LogRecordPos {
	art.lock.Lock()
	oldValue, _ := art.tree.Insert(key, pos)
	art.lock.Unlock()
	if oldValue == nil {
		return nil
	}
	return oldValue.(*data.LogRecordPos)
}

func (art *AdaptiveRadixTree) Get(key []byte) *data.LogRecordPos {
	art.lock.RLock()
	value, found := art.tree.Search(key)
	art.lock.RUnlock()
	if !found {
		return nil
	}
	return value.(*data.LogRecordPos)
}

func (art *AdaptiveRadixTree) Delete(key []byte) (*data.LogRecordPos, bool) {
	art.lock.Lock()
	oldValue, deleted := art.tree.Delete(key)
	art.lock.Unlock()
	if oldValue == nil {
		return nil, false
	}
	return oldValue.(*data.LogRecordPos), deleted
}

func (art *AdaptiveRadixTree) Size() int {
	art.lock.RLock()
	size := art.tree.Size()
	art.lock.RUnlock()
	return size
}

func (art *AdaptiveRadixTree) Iterator(reverse bool) Iterator {
	art.lock.Lock()
	defer art.lock.Unlock()
	return newARTIterator(art.tree, reverse)
}

func (art *AdaptiveRadixTree) Close() error {
	return nil
}

type artIterator struct {
	currIndex int     // 当前索引下标
	reverse   bool    // 是否是反向遍历
	values    []*Item // key+位置索引信息
}

func newARTIterator(tree goart.Tree, reverse bool) *artIterator {
	var idx int
	if reverse {
		idx = tree.Size() - 1
	}
	// 这里需要把数据都放到内存是因为我们的自定义的btree没有支持遍历的方法，若新索引有支持这不需要这样进行
	values := make([]*Item, tree.Size())
	saveValues := func(node goart.Node) bool {
		item := &Item{
			key: node.Key(),
			pos: node.Value().(*data.LogRecordPos),
		}
		values[idx] = item
		if reverse {
			idx--
		} else {
			idx++
		}
		return true
	}

	tree.ForEach(saveValues)

	return &artIterator{
		currIndex: 0,
		reverse:   reverse,
		values:    values,
	}
}

// Rewind 重新回到迭代器的起点
func (ai *artIterator) Rewind() {
	ai.currIndex = 0
}

// Seek 根据传入的key查找到第一个大于（或小于）等于的模板key，根据这个key开始遍历
func (ai *artIterator) Seek(key []byte) {
	if ai.reverse {
		// 二分查找法，标准库提供
		ai.currIndex = sort.Search(len(ai.values), func(i int) bool {
			return bytes.Compare(ai.values[i].key, key) <= 0
		})
	} else {
		ai.currIndex = sort.Search(len(ai.values), func(i int) bool {
			return bytes.Compare(ai.values[i].key, key) >= 0
		})
	}
}

// Next 跳转到下一个key
func (ai *artIterator) Next() {
	ai.currIndex += 1
}

// Valid 是否有效，即是否已经遍历完了所有的key，用于退出遍历
func (ai *artIterator) Valid() bool {
	return ai.currIndex < len(ai.values)
}

// Key 当前遍历位置的key数据
func (ai *artIterator) Key() []byte {
	return ai.values[ai.currIndex].key
}

// Value 当前遍历位置的Value数据
func (ai *artIterator) Value() *data.LogRecordPos {
	return ai.values[ai.currIndex].pos
}

// Close 关闭迭代器，释放相应资源
func (ai *artIterator) Close() {
	ai.values = nil
}
