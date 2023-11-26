package index

import (
	"bitcask-go/data"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestBPlusTree_Put(t *testing.T) {
	path := filepath.Join(os.TempDir(), "bptree-put")
	_ = os.MkdirAll(path, os.ModePerm)
	defer func() {
		_ = os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)
	res1 := tree.Put([]byte("aaa"), &data.LogRecordPos{Fid: 123, Offset: 456})
	assert.Nil(t, res1)
	res2 := tree.Put([]byte("bbb"), &data.LogRecordPos{Fid: 123, Offset: 456})
	assert.Nil(t, res2)
	res3 := tree.Put([]byte("ccc"), &data.LogRecordPos{Fid: 123, Offset: 456})
	assert.Nil(t, res3)

	res4 := tree.Put([]byte("aaa"), &data.LogRecordPos{Fid: 1234, Offset: 7456})
	assert.Equal(t, res4.Fid, uint32(123))
	assert.Equal(t, res4.Offset, int64(456))

}

func TestBPlusTree_Get(t *testing.T) {
	path := filepath.Join(os.TempDir(), "bptree-get")
	_ = os.MkdirAll(path, os.ModePerm)
	defer func() {
		_ = os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)

	pos := tree.Get([]byte("not exist"))
	t.Log(pos)
	assert.Nil(t, pos)

	tree.Put([]byte("aaa"), &data.LogRecordPos{Fid: 123, Offset: 456})
	pos1 := tree.Get([]byte("aaa"))
	t.Log(pos1)
	assert.NotNil(t, pos1)

	tree.Put([]byte("aaa"), &data.LogRecordPos{Fid: 321, Offset: 456})
	pos2 := tree.Get([]byte("aaa"))
	t.Log(pos2)
	assert.NotNil(t, pos2)
}

func TestBPlusTree_Delete(t *testing.T) {
	path := filepath.Join(os.TempDir(), "bptree-delete")
	_ = os.MkdirAll(path, os.ModePerm)
	defer func() {
		_ = os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)

	res1, ok1 := tree.Delete([]byte("not exist"))
	assert.False(t, ok1)
	assert.Nil(t, res1)

	tree.Put([]byte("aaa"), &data.LogRecordPos{Fid: 123, Offset: 456})
	res2, ok2 := tree.Delete([]byte("aaa"))
	assert.True(t, ok2)
	assert.Equal(t, res2.Fid, uint32(123))
	assert.Equal(t, res2.Offset, int64(456))

	pos1 := tree.Get([]byte("aaa"))
	t.Log(pos1)
	assert.Nil(t, pos1)
}

func TestBPlusTree_Size(t *testing.T) {
	path := filepath.Join(os.TempDir(), "bptree-size")
	_ = os.MkdirAll(path, os.ModePerm)
	defer func() {
		_ = os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)

	assert.Equal(t, 0, tree.Size())

	tree.Put([]byte("aac"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("abc"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("acc"), &data.LogRecordPos{Fid: 123, Offset: 999})

	assert.Equal(t, 3, tree.Size())
}

func TestBPlusTree_Iterator(t *testing.T) {
	path := filepath.Join(os.TempDir(), "bptree-iter")
	_ = os.MkdirAll(path, os.ModePerm)
	defer func() {
		_ = os.RemoveAll(path)
	}()
	tree := NewBPlusTree(path, false)

	tree.Put([]byte("aaa"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("bbb"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("ccc"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("ddd"), &data.LogRecordPos{Fid: 123, Offset: 999})
	tree.Put([]byte("eee"), &data.LogRecordPos{Fid: 123, Offset: 999})

	iter := tree.Iterator(true)
	for iter.Rewind(); iter.Valid(); iter.Next() {
		assert.NotNil(t, iter.Key())
		assert.NotNil(t, iter.Value())
	}
}
