package bitmap

import "testing"

func TestBM(t *testing.T) {
	bm := New(511)

	// all should be unset
	for i := 0; i < 511; i++ {
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

		bm.SetAtomic(i, false)
		if bm.Get(i) {
			t.Errorf("bit %d is set while it shouldn't", i)
		}
	}
}
