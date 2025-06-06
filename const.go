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

var SimpleFontMap = map[rune][][]int{
	'0': {
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
	},
	'1': {
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
	},
	'2': {
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	},
	'3': {
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
	},
	'4': {
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 1, 1, 0, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0, 1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0, 1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
	},
	'5': {
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
	},
	'6': {
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
	},
	'7': {
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	'8': {
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
	},
	'9': {
		{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
	},
}
