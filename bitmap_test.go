package bitmap

import (
	"sync"
	"testing"
)

func TestBM(t *testing.T) {
	bm := New(511)

	// all should be unset
	for i := 0; i <= 511; i++ {
		if bm.Get(i) {
			t.Errorf("bit %d is set while it shouldn't (initial)", i)
		}
		if bm.GetAtomic(i) {
			t.Errorf("bit %d is set while it shouldn't (initial, atomic)", i)
		}
	}

	v := []int{5, 7, 8, 9, 15, 300, 350, 255, 256, 257, 258, 259, 280}

	for i := range v {
		bm.Set(i, true)

		if !bm.Get(i) {
			t.Errorf("bit %d isn't set while it should", i)
		}
		if bm.Get(i - 1) {
			t.Errorf("bit %d-1 is set while it shouldn't", i)
		}
		if bm.Get(i + 1) {
			t.Errorf("bit %d+1 is set while it shouldn't", i)
		}

		// test atomic
		if !bm.GetAtomic(i) {
			t.Errorf("bit %d isn't set while it should (atomic)", i)
		}

		switch i % 3 {
		case 0:
			bm.SetAtomic(i, false)
		case 1:
			bm.Set(i, false)
		case 2:
			bm.Toggle(i)
		}

		if bm.Get(i) {
			t.Errorf("bit %d is set while it shouldn't", i)
		}
	}

	// atomic test
	var wg sync.WaitGroup
	wg.Add(512)

	var lk sync.RWMutex
	lk.Lock()

	for i := 0; i <= 511; i++ {
		go func(j int) {
			lk.RLock()
			defer lk.RUnlock()

			if j%2 == 0 {
				bm.SetAtomic(j, true)
			} else {
				bm.ToggleAtomic(j)
			}
			wg.Done()
		}(i)
	}

	lk.Unlock() // release all the threads waiting for rlock
	wg.Wait()

	// all should be set
	for i := 0; i <= 511; i++ {
		if !bm.Get(i) {
			t.Errorf("bit %d is not set while it should (final)", i)
		}
		if !bm.GetAtomic(i) {
			t.Errorf("bit %d is not set while it should (final, atomic)", i)
		}
	}
}
