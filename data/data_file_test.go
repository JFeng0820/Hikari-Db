package data

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOpenDataFile(t *testing.T) {
	dataFile1, err := OpenDataFile(os.TempDir(), 0)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile1)

	dataFile2, err := OpenDataFile(os.TempDir(), 111)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile2)
}

func TestDataFile_Write(t *testing.T) {
	dataFile, err := OpenDataFile(os.TempDir(), 0)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile)

	err = dataFile.Write([]byte("a"))
	assert.Nil(t, err)

	err = dataFile.Write([]byte("b"))
	assert.Nil(t, err)

	err = dataFile.Write([]byte("c"))
	assert.Nil(t, err)
}

func TestDataFile_Close(t *testing.T) {
	dataFile1, err := OpenDataFile(os.TempDir(), 123)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile1)

	err = dataFile1.Write([]byte("a"))
	assert.Nil(t, err)

	err = dataFile1.Close()
	assert.Nil(t, err)
}

func TestDataFile_Sync(t *testing.T) {
	dataFile1, err := OpenDataFile(os.TempDir(), 123)
	assert.Nil(t, err)
	assert.NotNil(t, dataFile1)

	err = dataFile1.Write([]byte("a"))
	assert.Nil(t, err)

	err = dataFile1.Sync()
	assert.Nil(t, err)
}
