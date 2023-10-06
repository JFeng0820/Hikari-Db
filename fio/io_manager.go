package fio

const DataFilePerm = 0644

// IOManager 通用 IO 管理接口，可以接入不同的 IO 类型
type IOManager interface {
	// Read 从文件的给定位置读取对应的数据
	Read([]byte, int64) (int, error)

	// Writer 写入字节数组到文件中
	Writer([]byte) (int, error)

	// Sync 持久化数据
	Sync() error

	// Close 关闭文件
	Close() error
}
