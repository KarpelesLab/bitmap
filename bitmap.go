package bitmap

import (
	"sync/atomic"
	"unsafe"
)

type Bitmap []byte

// New will spawn a new bitmap large enough
func New(size int) Bitmap {
	bytes := size / 8
	if size%8 != 0 {
		bytes += 1
	}

	// rem is used to allocate extra bytes for atomic operations by rounding capacity to next 32 bits
	rem := 3 - (size % 4)
	return make([]byte, size, size+rem)
}

func (m Bitmap) Get(i int) bool {
	return m[i/8]&(1<<byte(7-(i%8))) != 0
}

func (m Bitmap) Set(i int, v bool) {
	mask := byte(1 << byte(7-(i%8)))

	if v {
		m[i/8] |= mask
	} else {
		m[i/8] &= ^mask
	}
}

func (m Bitmap) Toggle(i int) {
	mask := byte(1 << byte(7-(i%8)))

	m[i/8] ^= mask
}

func (m Bitmap) GetAtomic(i int) bool {
	// get bit based on atomic operation
	u32 := (*uint32)(unsafe.Pointer(&m[int(i/32)*4]))
	v := atomic.LoadUint32(u32)

	return v&bitMask32(i) != 0
}

func (m Bitmap) SetAtomic(i int, v bool) {
	// set bit based on atomic operation
	u32 := (*uint32)(unsafe.Pointer(&m[int(i/32)*4]))

	mask := bitMask32(i)

	var n uint32
	for {
		old := atomic.LoadUint32(u32)

		if v {
			n = old | mask
		} else {
			n = old & ^mask
		}

		if old == n {
			// already set or unset
			return
		}

		if atomic.CompareAndSwapUint32(u32, old, n) {
			return
		}
	}
}

func (m Bitmap) ToggleAtomic(i int) {
	// set bit based on atomic operation
	u32 := (*uint32)(unsafe.Pointer(&m[int(i/32)*4]))

	mask := bitMask32(i)

	var n uint32
	for {
		old := atomic.LoadUint32(u32)

		n = old ^ mask

		if atomic.CompareAndSwapUint32(u32, old, n) {
			return
		}
	}
}
