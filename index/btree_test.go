package index

import (
	"github.com/stretchr/testify/assert"
	"kv-db/bitcask/data"
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
