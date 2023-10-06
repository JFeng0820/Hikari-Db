package data

type LogRecordType = byte

const (
	LogRecordNormal LogRecordType = iota
	LogRecordDeleted
)

// LogRecord 写入到数据文件的记录，数据文件中的数据是追加写入的，类似日志的格式
type LogRecord struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}

// LogRecordPos 数据内存索引， 主要描述数据在磁盘上的位置
type LogRecordPos struct {
	Fid    uint32 // 文件 id, 表示将数据存储到拿个文件
	Offset int64  // 偏移量
}

// EncodeLogRecord 对 LogRecord 进行编码， 返回字节数组以及长度
func EncodeLogRecord(logRecord *LogRecord) ([]byte, int64) {
	return nil, 0
}
