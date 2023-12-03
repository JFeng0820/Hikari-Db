package utils

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
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

// CopyDir 拷贝数据目录
func CopyDir(src, dest string, exclude []string) error {
	// 目标目录不存在则创建
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			return err
		}
	}

	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		fileName := strings.Replace(path, src, "", 1)
		if fileName == "" {
			return nil
		}

		for _, e := range exclude {
			matched, err := filepath.Match(e, info.Name())
			if err != nil {
				return err
			}
			if matched {
				return nil
			}
		}

		if info.IsDir() {
			return os.MkdirAll(filepath.Join(dest, fileName), info.Mode())
		}

		data, err := os.ReadFile(filepath.Join(src, fileName))
		if err != nil {
			return err
		}

		return os.WriteFile(filepath.Join(dest, fileName), data, info.Mode())
	})
}

func GetFileBase(filePath string) string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Base(filePath)
	default:
		return path.Base(filePath)
	}
}
