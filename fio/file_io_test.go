package fio

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func destroyFile(name string) {
	if err := os.RemoveAll(name); err != nil {
		panic(err)
	}
}

func TestNewFileIOManager(t *testing.T) {
	path := filepath.Join("/tmp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, fio)
}

func TestFileIO_Writer(t *testing.T) {
	path := filepath.Join("/tmp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	n, err := fio.Writer([]byte(""))
	assert.Equal(t, 0, n)
	assert.Nil(t, err)

	n, err = fio.Writer([]byte("hello"))
	assert.Equal(t, 5, n)
	assert.Nil(t, err)

	n, err = fio.Writer([]byte("world"))
	assert.Equal(t, 5, n)
	assert.Nil(t, err)

}

func TestFileIO_Read(t *testing.T) {
	path := filepath.Join("/tmp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	_, err = fio.Writer([]byte("key-a"))
	assert.Nil(t, err)

	_, err = fio.Writer([]byte("key-b"))
	assert.Nil(t, err)

	b := make([]byte, 5)
	n, err := fio.Read(b, 0)
	t.Log(string(b), n)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-a"), b)

	b1 := make([]byte, 5)
	n, err = fio.Read(b1, 5)
	t.Log(string(b1), n)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-b"), b1)
}

func TestFileIO_Sync(t *testing.T) {
	path := filepath.Join("/tmp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Sync()
	assert.Nil(t, err)
}

func TestFileIO_Close(t *testing.T) {
	path := filepath.Join("/tmp", "a.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Close()
	assert.Nil(t, err)
}
