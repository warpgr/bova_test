package store

import "github.com/warpgr/bova_test/pkg/store/internal"

type KVStorage[KT comparable, VT any] interface {
	Load(key KT) (VT, error)
	Store(key KT, value VT) error
	LoadMany(keys []KT) (map[KT]VT, error)
	StoreMany(elements map[KT]VT) error
	LoadAll() map[KT]VT
}

func NewKVMapStorage[KT comparable, VT any](initialSize int) KVStorage[KT, VT] {
	return internal.NewSafeMap[KT, VT](initialSize)
}
