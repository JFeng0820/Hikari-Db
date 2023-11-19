package index

import (
	"bitcask-go/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdaptiveRadixTree_Put(t *testing.T) {
	art := NewART()
	art.Put([]byte("key-1"), &data.LogRecordPos{Fid: 1, Offset: 12})
	art.Put([]byte("key-2"), &data.LogRecordPos{Fid: 1, Offset: 12})
	art.Put([]byte("key-3"), &data.LogRecordPos{Fid: 1, Offset: 12})
}

func TestAdaptiveRadixTree_Get(t *testing.T) {
	art := NewART()
	art.Put([]byte("key-1"), &data.LogRecordPos{Fid: 1, Offset: 12})
	pos := art.Get([]byte("key-1"))
	t.Log(pos)
	assert.NotNil(t, pos)

	pos1 := art.Get([]byte("not exsit"))
	assert.Nil(t, pos1)

	art.Put([]byte("key-1"), &data.LogRecordPos{Fid: 1123, Offset: 998})
	pos2 := art.Get([]byte("key-1"))
	t.Log(pos2)
	assert.NotNil(t, pos2)
}

func TestAdaptiveRadixTree_Delete(t *testing.T) {
	art := NewART()
	res := art.Delete([]byte("not exist"))
	assert.False(t, res)

	art.Put([]byte("key-1"), &data.LogRecordPos{Fid: 1, Offset: 12})
	pos1 := art.Get([]byte("key-1"))
	t.Log(pos1)
	assert.NotNil(t, pos1)
	res2 := art.Delete([]byte("key-1"))
	assert.True(t, res2)
	pos2 := art.Get([]byte("key-1"))
	t.Log(pos2)
	assert.Nil(t, pos2)
}

func TestAdaptiveRadixTree_Iterator(t *testing.T) {
	art := NewART()
	art.Put([]byte("key-1"), &data.LogRecordPos{Fid: 1, Offset: 12})
	art.Put([]byte("key-2"), &data.LogRecordPos{Fid: 1, Offset: 12})
	art.Put([]byte("key-3"), &data.LogRecordPos{Fid: 1, Offset: 12})
	art.Put([]byte("key-4"), &data.LogRecordPos{Fid: 1, Offset: 12})

	iter := art.Iterator(true)
	for iter.Rewind(); iter.Valid(); iter.Next() {
		t.Log(string(iter.Key()))
	}
	t.Log("=================")
	iter1 := art.Iterator(false)
	for iter1.Rewind(); iter1.Valid(); iter1.Next() {
		t.Log(string(iter1.Key()))
	}
}
