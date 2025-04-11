package gotool

var (
	// Mac 地址
	Mac string
	// EmptyMapBytes 空map
	EmptyMapBytes = []byte{0x7b, 0x7d}
	// DateTimeMilli 毫秒格式时间
	DateTimeMilli = "2006-01-02 15:04:05.000"
)

var ImageExtMap = map[string]struct{}{
	".jpg":  {},
	".jpeg": {},
	".png":  {},
	".gif":  {},
	".bmp":  {},
	".webp": {},
	".tif":  {},
	".tiff": {},
}
