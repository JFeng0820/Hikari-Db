package redis

import (
	bitcask "bitcask-go"
	"bitcask-go/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestRedisDataStructure_Get(t *testing.T) {
	opts := bitcask.DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go-redis-get")
	opts.DirPath = dir
	rds, err := NewRedisDataStructure(opts)
	assert.Nil(t, err)

	err = rds.Set(utils.GetTestKey(1), 0, utils.RandomValue(100))
	assert.Nil(t, err)
	err = rds.Set(utils.GetTestKey(2), time.Second*5, utils.RandomValue(100))
	assert.Nil(t, err)

	val1, err := rds.Get(utils.GetTestKey(1))
	assert.Nil(t, err)
	assert.NotNil(t, val1)

	val2, err := rds.Get(utils.GetTestKey(2))
	assert.Nil(t, err)
	assert.NotNil(t, val2)

	_, err = rds.Get(utils.GetTestKey(33))
	assert.Equal(t, bitcask.ErrKeyNotFound, err)
}

func TestRedisDataStructure_Del(t *testing.T) {
	opts := bitcask.DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go-redis-del")
	opts.DirPath = dir
	rds, err := NewRedisDataStructure(opts)
	assert.Nil(t, err)

	// del
	err = rds.Del(utils.GetTestKey(11))
	assert.Nil(t, err)

	err = rds.Set(utils.GetTestKey(1), 0, utils.RandomValue(100))
	assert.Nil(t, err)
	err = rds.Del(utils.GetTestKey(1))
	assert.Nil(t, err)

	_, err = rds.Get(utils.GetTestKey(1))
	assert.Equal(t, bitcask.ErrKeyNotFound, err)
}

func TestRedisDataStructure_Type(t *testing.T) {
	opts := bitcask.DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go-redis-type")
	opts.DirPath = dir
	rds, err := NewRedisDataStructure(opts)
	assert.Nil(t, err)

	err = rds.Set(utils.GetTestKey(1), 0, utils.RandomValue(100))
	assert.Nil(t, err)

	typ, err := rds.Type(utils.GetTestKey(1))
	assert.Nil(t, err)
	assert.Equal(t, String, typ)
}

func TestRedisDataStructure_HGet(t *testing.T) {
	opts := bitcask.DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go-redis-hget")
	opts.DirPath = dir
	rds, err := NewRedisDataStructure(opts)
	assert.Nil(t, err)

	ok1, err := rds.HSet(utils.GetTestKey(1), []byte("aa"), utils.RandomValue(100))
	assert.Nil(t, err)
	assert.True(t, ok1)

	v1 := utils.RandomValue(100)
	ok2, err := rds.HSet(utils.GetTestKey(1), []byte("aa"), v1)
	assert.Nil(t, err)
	assert.False(t, ok2)

	v2 := utils.RandomValue(100)
	ok3, err := rds.HSet(utils.GetTestKey(1), []byte("v"), v2)
	assert.Nil(t, err)
	assert.True(t, ok3)

	val1, err := rds.HGet(utils.GetTestKey(1), []byte("aa"))
	assert.Nil(t, err)
	assert.Equal(t, v1, val1)

	val2, err := rds.HGet(utils.GetTestKey(1), []byte("v"))
	assert.Nil(t, err)
	assert.Equal(t, v2, val2)

	_, err = rds.HGet(utils.GetTestKey(1), []byte("field-not-exist"))
	assert.Equal(t, bitcask.ErrKeyNotFound, err)
}
