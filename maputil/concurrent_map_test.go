package maputil

import (
	"github.com/up-zero/gotool/testutil"
	"strconv"
	"sync"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	// 示例 1: 基本操作 (string -> int)
	m := NewConcurrentMap[string, int]()

	// Set 并发写
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := "key_" + strconv.Itoa(n)
			m.Set(key, n)
		}(i)
	}
	wg.Wait() // 等待所有写入完成

	// Len
	testutil.Equal(t, m.Len(), 100)

	// Get 并发读
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			defer wg.Done()
			val, ok := m.Get("key_" + strconv.Itoa(n))
			if ok {
				testutil.Equal(t, val, n)
			}
		}(i)
	}
	wg.Wait()

	// Get 获取不存在的
	_, ok := m.Get("key_not_exist")
	testutil.Equal(t, ok, false)

	// GetOrSet
	// 尝试获取 "key_10"，它已存在
	val, loaded := m.GetOrSet("key_10", 999)
	testutil.Equal(t, val, 10)
	testutil.Equal(t, loaded, true)
	// 尝试获取 "new_key"，它不存在
	val, loaded = m.GetOrSet("new_key", 999)
	testutil.Equal(t, val, 999)
	testutil.Equal(t, loaded, false)

	// Range 迭代
	count := 0
	m.Range(func(key string, value int) bool {
		testutil.Equal(t, key, "key_"+strconv.Itoa(value))
		count++
		return count < 3 // 只打印3个
	})

	// Delete
	m.Delete("key_99")
	_, ok = m.Get("key_99")
	testutil.Equal(t, ok, false)

	// Clear
	m.Clear()
	testutil.Equal(t, m.Len(), 0)

	// 示例 2: 自定义结构体 (int -> User)
	userMap := NewConcurrentMap[int, User]()

	userMap.Set(1, User{ID: 1, Name: "Alice"})
	userMap.Set(2, User{ID: 2, Name: "Bob"})

	alice, _ := userMap.Get(1)
	testutil.Equal(t, alice.ID, 1)
}
