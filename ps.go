package gotool

type Process struct {
	Pid  string `json:"pid"`  // 进程ID号
	PPid string `json:"ppid"` // 进程父级ID号
	Cmd  string `json:"cmd"`  // 执行的命令
}
