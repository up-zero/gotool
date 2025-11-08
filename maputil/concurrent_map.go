package maputil

import "sync"

// ConcurrentMap 是一个类型安全、并发安全的 Map 封装
// 它使用泛型来保证 K (键) 和 V (值) 的类型，
// 并通过 sync.RWMutex 来保证并发安全
type ConcurrentMap[K comparable, V any] struct {
	mu    sync.RWMutex
	items map[K]V
}

// NewConcurrentMap 初始化并发 Map
func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		items: make(map[K]V),
	}
}

// Set 设置键值对（写操作）
func (m *ConcurrentMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[key] = value
}

// Get 获取值（读操作）
func (m *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok := m.items[key]

	return value, ok
}

// GetOrSet 获取一个键的值（读操作），如果键不存在，则设置并返回新值（写操作）
//
//   - loaded 为 true 表示值是存在的
//   - loaded 为 false 表示值是新设置的
func (m *ConcurrentMap[K, V]) GetOrSet(key K, value V) (actual V, loaded bool) {
	// 尝试快速读取（读锁）
	m.mu.RLock()
	actual, loaded = m.items[key]
	m.mu.RUnlock()
	if loaded {
		return actual, true
	}

	// 值不存在，升级为写锁
	m.mu.Lock()
	defer m.mu.Unlock()

	// 双重检查锁定
	// 在我们等待写锁期间，可能有其他 goroutine 已经设置了值
	actual, loaded = m.items[key]
	if loaded {
		return actual, true
	}

	m.items[key] = value
	return value, false
}

// Delete 删除键（写操作）
func (m *ConcurrentMap[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.items, key)
}

// GetAndDelete 获取键的值并删除（写操作）
func (m *ConcurrentMap[K, V]) GetAndDelete(key K) (V, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	value, ok := m.items[key]
	if ok {
		delete(m.items, key)
	}
	return value, ok
}

// Len 获取 Map 中元素的数量（读操作）
func (m *ConcurrentMap[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.items)
}

// Clear 清空 Map 中的所有元素（写操作）
func (m *ConcurrentMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	// 为了最快地清空并释放内存，直接创建一个新 map
	m.items = make(map[K]V)
}

// Range 遍历 Map（读操作）
//
//	对每个键值对执行 f 函数，如果 f 返回 false，则停止遍历
//
//	注意：遍历期间会持有读锁，因此请勿在 f 中执行
//	任何可能导致死锁（如调用 m.Set/m.Delete）或长时间阻塞的操作。
func (m *ConcurrentMap[K, V]) Range(f func(key K, value V) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for key, value := range m.items {
		if !f(key, value) {
			break
		}
	}
}
