package bitcask

import (
	"bitcask-go/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDB_NewIterator(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go-1")
	opts.DirPath = dir
	db, err := Open(opts)
	defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	iterator := db.NewIterator(DefaultIteratorOptions)
	defer iterator.Close()
	assert.NotNil(t, iterator)
	assert.Equal(t, false, iterator.Valid())
}

func TestDB_NewIterator_One_Value(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go-2")
	opts.DirPath = dir
	db, err := Open(opts)
	defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.Put(utils.GetTestKey(10), utils.GetTestKey(10))
	assert.Nil(t, err)

	iterator := db.NewIterator(DefaultIteratorOptions)
	defer iterator.Close()
	assert.NotNil(t, iterator)
	assert.Equal(t, true, iterator.Valid())
	assert.Equal(t, utils.GetTestKey(10), iterator.Key())
	val, err := iterator.Value()
	assert.Nil(t, err)
	assert.Equal(t, utils.GetTestKey(10), val)
}

func TestDB_NewIterator_Multi_Value(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go-3")
	opts.DirPath = dir
	db, err := Open(opts)
	defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.Put([]byte("andne"), utils.RandomValue(10))
	assert.Nil(t, err)
	err = db.Put([]byte("and32"), utils.RandomValue(10))
	assert.Nil(t, err)
	err = db.Put([]byte("andqw"), utils.RandomValue(10))
	assert.Nil(t, err)
	err = db.Put([]byte("ewqe"), utils.RandomValue(10))
	assert.Nil(t, err)
	err = db.Put([]byte("rsasd"), utils.RandomValue(10))
	assert.Nil(t, err)

	t.Log("============ 正向迭代start ============")

	iterator := db.NewIterator(DefaultIteratorOptions)
	defer iterator.Close()
	for iterator.Rewind(); iterator.Valid(); iterator.Next() {
		t.Log("key = ", string(iterator.Key()))
		assert.NotNil(t, iterator.Key())
	}
	t.Log("============ 正向迭代end ============")
	t.Log("============ 正向迭代 seek start ============")
	iterator.Rewind()
	for iterator.Seek([]byte("b")); iterator.Valid(); iterator.Next() {
		t.Log("key = ", string(iterator.Key()))
		assert.NotNil(t, iterator.Key())
	}
	t.Log("============ 正向迭代 seek end ============")

	t.Log("============ 反向迭代 start ============")
	iteratorRev := DefaultIteratorOptions
	iteratorRev.Reverse = true
	iterator2 := db.NewIterator(iteratorRev)
	defer iterator.Close()
	for iterator2.Rewind(); iterator2.Valid(); iterator2.Next() {
		t.Log("key = ", string(iterator2.Key()))
		assert.NotNil(t, iterator2.Key())
	}
	t.Log("============ 反向迭代 end ============")
	t.Log("============ 反向迭代 seek start ============")
	iterator2.Rewind()
	for iterator2.Seek([]byte("b")); iterator2.Valid(); iterator2.Next() {
		t.Log("key = ", string(iterator2.Key()))
		assert.NotNil(t, iterator2.Key())
	}
	t.Log("============ 反向迭代 seek end ============")

	t.Log("============ 正向迭代 prefix start ============")
	iteratorOption3 := DefaultIteratorOptions
	iteratorOption3.Prefix = []byte("a")
	iterator3 := db.NewIterator(iteratorOption3)
	defer iterator.Close()
	for iterator3.Rewind(); iterator3.Valid(); iterator3.Next() {
		t.Log("key = ", string(iterator3.Key()))
		assert.NotNil(t, iterator3.Key())
	}
	t.Log("============ 正向迭代 prefix end ============")

	t.Log("============ 逆向迭代 prefix start ============")
	iteratorOption4 := DefaultIteratorOptions
	iteratorOption4.Prefix = []byte("a")
	iteratorOption4.Reverse = true
	iterator4 := db.NewIterator(iteratorOption4)
	defer iterator.Close()
	for iterator4.Rewind(); iterator4.Valid(); iterator4.Next() {
		t.Log("key = ", string(iterator4.Key()))
		assert.NotNil(t, iterator4.Key())
	}
	t.Log("============ 逆向迭代 prefix end ============")

}
