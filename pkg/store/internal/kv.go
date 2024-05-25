package internal

import (
	"sync"

	"github.com/warpgr/bova_test/pkg/exchanges/common"
)

type SafeMap[KT comparable, VT any] struct {
	store map[KT]VT
	guard sync.RWMutex
}

func NewSafeMap[KT comparable, VT any](size int) *SafeMap[KT, VT] {
	return &SafeMap[KT, VT]{
		store: make(map[KT]VT, size),
		guard: sync.RWMutex{},
	}
}

func (sm *SafeMap[KT, VT]) Load(key KT) (VT, error) {
	sm.guard.RLock()
	defer sm.guard.RUnlock()
	value, exists := sm.store[key]
	if !exists {
		return value, common.ErrLoadElement
	}
	return value, nil
}

func (sm *SafeMap[KT, VT]) Store(key KT, value VT) error {
	sm.guard.Lock()
	defer sm.guard.Unlock()

	sm.store[key] = value
	return nil
}

func (sm *SafeMap[KT, VT]) LoadMany(keys []KT) (map[KT]VT, error) {
	sm.guard.RLock()
	defer sm.guard.RUnlock()

	elements := make(map[KT]VT, len(keys))
	for _, key := range keys {
		element, exists := sm.store[key]
		if exists {
			elements[key] = element
		}
	}
	return elements, nil
}

func (sm *SafeMap[KT, VT]) StoreMany(elements map[KT]VT) error {
	sm.guard.Lock()
	defer sm.guard.Unlock()

	for key, value := range elements {
		sm.store[key] = value
	}
	return nil
}

func (sm *SafeMap[KT, VT]) LoadAll() map[KT]VT {
	sm.guard.RLock()
	defer sm.guard.RUnlock()

	elements := make(map[KT]VT, len(sm.store))
	for key, value := range sm.store {
		elements[key] = value
	}

	return elements
}
