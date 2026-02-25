package fileutil

type UnzipNotify struct {
	Progress int    `json:"progress"` // 解压进度
	FPath    string `json:"f_path"`   // 解压文件的路径
}

// ReadLineCallback 是按行读取的回调函数定义
//
// # Params:
//
//	lineNum: 当前行号 (从 1 开始)
//	line: 当前行的文本内容
//
// # Returns:
//
//	bool: 如果返回 false，则立即终止后续读取；返回 true 则继续往下读
type ReadLineCallback func(lineNum int, line string) bool
