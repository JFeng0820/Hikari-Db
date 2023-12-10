//go:build linux

package linux_utils

import "syscall"

type LinuxDiskUsage struct{}

func (d LinuxDiskUsage) GetDiskUsage(filePath string) (uint64, error) {
	var stat syscall.Statfs_t
	if err = syscall.Statfs(wd, &stat); err != nil {
		return 0, err
	}
	// 计算非特权用户可用的空闲空间
	return stat.Bavail * uint64(stat.Bsize), nil
}
