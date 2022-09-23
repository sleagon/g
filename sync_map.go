package g

import "sync"

type SyncMap[K comparable, V any] struct {
	data sync.Map
}

func (m *SyncMap[K, V]) Load(key K) (V, bool) {
	r, ok := m.data.Load(key)
	if ok {
		return r.(V), true
	}
	var rv V
	return rv, false
}
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.data.Store(key, value)
}
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	r, ok := m.data.LoadOrStore(key, value)
	return r.(V), ok
}
func (m *SyncMap[K, V]) LoadAndDelete(key K) (V, bool) {
	r, ok := m.data.LoadAndDelete(key)
	if v1, ok1 := r.(V); ok1 {
		return v1, ok
	}
	var rv V
	return rv, false
}
func (m *SyncMap[K, V]) Delete(key K) {
	m.data.Delete(key)
}
func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.data.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func NewSyncMap[K comparable, V any]() SyncMap[K, V] {
	return SyncMap[K, V]{
		data: sync.Map{},
	}
}
