package utils

import (
	"io/fs"
	"path/filepath"
)

// DirSize 获取一个目录的大小
func DirSize(dirPath string) (int64, error) {
	var size int64
	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

// AvailableDiskSize 获取磁盘剩余可用空间大小
func AvailableDiskSize() (uint64, error) {
	// TODO 之后进行补充
	//wd, err := syscall.Getwd()
	//if err != nil {
	//	return 0, err
	//}
	//var stat unix.
	//if err = sys.Statfs(wd, &stat); err != nil {
	//	return 0, err
	//}
	//return stat.Bavail * uint64(stat.Bsize), nil
	return 0, nil
}
