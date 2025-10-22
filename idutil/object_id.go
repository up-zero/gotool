package idutil

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

// ObjectID 是一个 12 字节的唯一标识符。
type ObjectID [12]byte

var (
	objectIdOnce sync.Once
	// 主机随机部分（5 字节）
	objectIdHostRand [5]byte
	// 原子计数器（3 字节）
	objectIdCounter uint32
)

// NewObjectID 生成一个新的 ObjectID
func NewObjectID() ObjectID {
	objectIdOnce.Do(func() {
		// host random
		if _, err := io.ReadFull(rand.Reader, objectIdHostRand[:]); err != nil {
			panic(fmt.Errorf("cannot genera host rand: %w", err))
		}
		// counter
		var seed [4]byte
		if _, err := io.ReadFull(rand.Reader, seed[:]); err != nil {
			panic(fmt.Errorf("cannot genera counter: %w", err))
		}
		objectIdCounter = binary.BigEndian.Uint32(seed[:])
	})

	var b [12]byte

	// 当前时间戳秒（4 字节）
	ts := uint32(time.Now().Unix())
	binary.BigEndian.PutUint32(b[0:4], ts)

	// 主机随机部分（5 字节）
	copy(b[4:9], objectIdHostRand[:])

	// 原子计数器（3 字节）
	c := atomic.AddUint32(&objectIdCounter, 1)
	b[9] = byte(c >> 16)
	b[10] = byte(c >> 8)
	b[11] = byte(c)

	return b
}

// String 返回 24 位十六进制字符串（如 "68f5ee80e208bcd1d3eacec9"）
func (oid ObjectID) String() string {
	return hex.EncodeToString(oid[:])
}

// Timestamp 从 ObjectID 中提取生成时间
func (oid ObjectID) Timestamp() time.Time {
	ts := binary.BigEndian.Uint32(oid[0:4])
	return time.Unix(int64(ts), 0)
}
