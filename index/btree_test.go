package index

import (
	"bitcask-go/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBtree_Put(t *testing.T) {
	bt := NewBTree()
	res1 := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res1)

	res2 := bt.Put([]byte("key"), &data.LogRecordPos{Fid: 1, Offset: 2})
	assert.True(t, res2)
}

func TestBtree_Get(t *testing.T) {
	bt := NewBTree()
	res1 := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res1)

	po1 := bt.Get(nil)
	assert.Equal(t, uint32(1), po1.Fid)
	assert.Equal(t, int64(100), po1.Offset)

	res2 := bt.Put([]byte("key"), &data.LogRecordPos{Fid: 1, Offset: 2})
	assert.True(t, res2)
	res3 := bt.Put([]byte("key"), &data.LogRecordPos{Fid: 2, Offset: 3})
	assert.True(t, res3)

	po2 := bt.Get([]byte("key"))
	assert.Equal(t, uint32(2), po2.Fid)
	assert.Equal(t, int64(3), po2.Offset)
	t.Log(po2)
}

func TestBtree_Delete(t *testing.T) {
	bt := NewBTree()
	res1 := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res1)
	res2 := bt.Delete(nil)
	assert.True(t, res2)

	res3 := bt.Put([]byte("key"), &data.LogRecordPos{Fid: 22, Offset: 33})
	assert.True(t, res3)
	res4 := bt.Delete([]byte("key"))
	assert.True(t, res4)
}

func TestBtree_Iterator(t *testing.T) {
	bt1 := NewBTree()
	// 1.Btree 为空的情况
	iter1 := bt1.Iterator(false)
	assert.Equal(t, false, iter1.Valid())

	// 2.Btree 有数据的情况下
	bt1.Put([]byte("code"), &data.LogRecordPos{Fid: 1, Offset: 10})
	iter2 := bt1.Iterator(false)
	assert.Equal(t, true, iter2.Valid())
	assert.NotNil(t, iter2.Value())
	assert.NotNil(t, iter2.Key())
	iter2.Next()
	assert.Equal(t, false, iter2.Valid())

	// 有多条数据
	bt1.Put([]byte("code"), &data.LogRecordPos{Fid: 1, Offset: 10})
	bt1.Put([]byte("eeee1"), &data.LogRecordPos{Fid: 1, Offset: 10})
	bt1.Put([]byte("fgsd"), &data.LogRecordPos{Fid: 1, Offset: 10})
	bt1.Put([]byte("fewr"), &data.LogRecordPos{Fid: 1, Offset: 10})
	t.Log("reverse False ===============")
	iter3 := bt1.Iterator(false)
	for iter3.Rewind(); iter3.Valid(); iter3.Next() {
		t.Log("key = ", string(iter3.Key()))
		assert.NotNil(t, iter3.Key())
	}

	t.Log("reverse True ===============")

	iter4 := bt1.Iterator(true)
	for iter4.Rewind(); iter4.Valid(); iter4.Next() {
		t.Log("key = ", string(iter4.Key()))
		assert.NotNil(t, iter4.Key())
	}

	t.Log("测试 seek reverse false ===============")

	// 4. 测试 seek
	iter5 := bt1.Iterator(false)
	for iter5.Seek([]byte("ee")); iter5.Valid(); iter5.Next() {
		t.Log(string(iter5.Key()))
		assert.NotNil(t, iter5.Key())
	}

	t.Log("测试 seek reverse true ===============")

	// 5. 测试 seek
	iter6 := bt1.Iterator(true)
	for iter6.Seek([]byte("ee")); iter6.Valid(); iter6.Next() {
		t.Log(string(iter6.Key()))
		assert.NotNil(t, iter6.Key())
	}

}
