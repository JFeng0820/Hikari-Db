//go:build windows

package window_utils

import (
	"fmt"
	"syscall"
	"unsafe"
)

type WindowsDiskUsage struct{}

func (d WindowsDiskUsage) GetDiskUsage(filePath string) (uint64, error) {
	var freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes uint64
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
	_, _, err := c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(filePath))),
		uintptr(unsafe.Pointer(&freeBytesAvailable)),
		uintptr(unsafe.Pointer(&totalNumberOfBytes)),
		uintptr(unsafe.Pointer(&totalNumberOfFreeBytes)),
	)
	if err != nil && err != syscall.Errno(0) {
		fmt.Printf("Error getting disk space: %s\n", err)
		return 0, err
	}
	return totalNumberOfFreeBytes, nil
}
