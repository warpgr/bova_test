package internal_test

import (
	"sync"
	"testing"

	"github.com/warpgr/bova_test/pkg/store/internal"

	"github.com/stretchr/testify/require"
)

func TestSafeMapStoreLoad(t *testing.T) {
	store := internal.NewSafeMap[int, int](10)

	ct := &counter{}
	storeJob := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		storeJob.Add(1)
		// Running concurrent storing jobs.
		go func() {
			defer storeJob.Done()
			for j := 0; j < 10; j++ {
				c := ct.next()
				err := store.Store(c, c*2)
				require.NoError(t, err)
			}
		}()
	}

	storeJob.Wait()
	for i := 1; i <= 100; i++ {
		v, err := store.Load(i)
		require.NoError(t, err)
		require.Equal(t, i*2, v)
	}
}

func TestSafeMapStoreLoadMany(t *testing.T) {
	store := internal.NewSafeMap[int, int](10)

	ct := &counter{}
	storeJob := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		storeJob.Add(1)
		// Running concurrent storing jobs.
		go func() {
			defer storeJob.Done()
			bunch := make(map[int]int, 10)
			for j := 0; j < 10; j++ {
				c := ct.next()
				bunch[c] = c * 2
			}
			err := store.StoreMany(bunch)
			require.NoError(t, err)
		}()
	}

	storeJob.Wait()
	for i := 1; i <= 90; i += 3 {
		vs, err := store.LoadMany([]int{i, i + 1, i + 2})
		require.NoError(t, err)
		require.Equal(t, i*2, vs[i])
		require.Equal(t, (i+1)*2, vs[i+1])
		require.Equal(t, (i+2)*2, vs[i+2])
	}

	vs := store.LoadAll()
	require.Equal(t, 100, len(vs))
}

type counter struct {
	c     int
	guard sync.Mutex
}

func (c *counter) next() int {
	c.guard.Lock()
	defer c.guard.Unlock()
	c.c++
	return c.c
}
