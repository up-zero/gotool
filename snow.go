package gotool

import (
	"sync"
	"time"
)

var snowflakeEntity *Snowflake

// Snowflake 结构体
type Snowflake struct {
	mutex      sync.Mutex
	epoch      int64
	nodeID     uint8
	sequence   uint16
	lastMillis int64
}

// NewSnowflake 创建 Snowflake 实例
//
// epoch: 起始时间戳
// nodeID: 节点ID
func NewSnowflake(epoch int64, nodeID uint8) *Snowflake {
	return &Snowflake{
		epoch:    epoch,
		nodeID:   nodeID,
		sequence: 0,
	}
}

// GenerateSnowflake 生成新的雪花码
func (s *Snowflake) GenerateSnowflake() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// 时钟回拨
	currentMillis := s.timeGen()
	if currentMillis < s.lastMillis {
		panic("Clock moved backwards. Refusing to generate id.")
	}

	// 相同毫秒内，序列号自增
	if currentMillis == s.lastMillis {
		s.sequence = (s.sequence + 1) & 0xFFF
		if s.sequence == 0 {
			currentMillis = s.tilNextMillis(s.lastMillis)
		}
	} else {
		// 不同毫秒内，序列号置为 0
		s.sequence = 0
	}
	s.lastMillis = currentMillis

	id := ((currentMillis - s.epoch) << 22) |
		(int64(s.nodeID) << 12) |
		int64(s.sequence)

	return id
}

// timeGen 获取当前时间戳
func (s *Snowflake) timeGen() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// tilNextMillis 获取下一个时间戳
func (s *Snowflake) tilNextMillis(lastMillis int64) int64 {
	millis := s.timeGen()
	for millis <= lastMillis {
		millis = s.timeGen()
	}
	return millis
}

// SignalSnowflake 单节点的雪花码
func SignalSnowflake() int64 {
	if snowflakeEntity == nil {
		snowflakeEntity = NewSnowflake(1600000000000, 1)
	}
	sf := snowflakeEntity
	return sf.GenerateSnowflake()
}
