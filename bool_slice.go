package sparse

import "sync"

// BoolSlice is a compact representation of a []bool.
type BoolSlice struct {
	m     []int64
	msize int64
	size  int64
	rw    sync.RWMutex
}

// NewBoolSlice returns a pointer on a new BoolSlice.
func NewBoolSlice() *BoolSlice {
	return &BoolSlice{
		m:     make([]int64, 1),
		msize: 1,
	}
}

// Size returns the size of the BoolSlice.
func (bs *BoolSlice) Size() (s int64) {
	bs.rw.RLock()
	s = bs.size
	bs.rw.RUnlock()
	return
}

// Get returns the bool at the given index. It returns false if the index is
// greater or equal to the return value of Size().
func (bs *BoolSlice) Get(idx int64) bool {
	bs.rw.RLock()
	defer bs.rw.RUnlock()
	if idx >= bs.size {
		return false
	}
	return bs.mvalue(bs.mindex(idx))
}

// Set sets a value at the given index. If the index is greater than the return
// value of Size() the missing indexes will be filled with false.
func (bs *BoolSlice) Set(idx int64, v bool) (err error) {
	//bs.rw.Lock()
	//defer bs.rw.Unlock()

	midx := bs.mindex(idx)

	// TODO test me
	if idx >= bs.size {
		lastval := bs.lastmvalue()

		if !v {
			if !lastval {
				// last value is false and the new value is false
				// -> simple increment
				bs.m[bs.msize-1] += idx - bs.size
				bs.size = idx + 1
			} else {
				// last value is true and the new value is false
				// -> append the missing false values
				bs.m = append(bs.m, idx-bs.size)
				bs.size = idx + 1
				bs.msize++
			}
		} else {
			if !lastval {
				// last value is false and the new value is true
				// -> increment + append true
				bs.m[bs.msize-1] += idx - bs.size
				bs.size = idx + 1
				bs.m = append(bs.m, 1)
				bs.msize++
			} else {
				// last value is true and the new value is true
				// -> append the missing false values + the true one
				bs.m = append(bs.m, idx-bs.size-1, 1)
				bs.msize += 2
			}
		}
		return
	}

	// noop
	if bs.mvalue(midx) == v {
		return
	}

	// FIXME this is too dirty
	s := bs.ToSlice()
	s[idx] = v
	*bs = *BoolSliceFromSlice(s)

	return
}

// Append appends a value to the slice.
func (bs *BoolSlice) Append(v bool) (err error) {
	bs.rw.Lock()
	defer bs.rw.Unlock()

	if bs.lastmvalue() == v {
		bs.m[bs.msize-1]++
	} else {
		bs.m = append(bs.m, 1)
		bs.msize++
	}
	bs.size++
	return
}

func (bs *BoolSlice) lastmvalue() bool {
	if bs.msize == 0 {
		return true
	}

	return bs.mvalue(bs.msize - 1)
}

func (bs *BoolSlice) mvalue(i int64) bool { return i&1 == 1 }

func (bs *BoolSlice) mindex(idx int64) int64 {
	var i, cursor int64

	for ; i < bs.msize; i++ {
		cursor += bs.m[i]
		if idx < cursor {
			return i
		}
	}

	return bs.msize
}

// BoolSliceFromSlice converts a []bool into a BoolSlice
func BoolSliceFromSlice(s []bool) *BoolSlice {
	var idx, msize int64
	var currval bool

	msize = 1

	slen := len(s)

	for i := 0; i < slen; i++ {
		if s[i] != currval {
			msize++
			currval = !currval
		}
	}

	bs := &BoolSlice{
		m:     make([]int64, msize),
		msize: msize,
		size:  int64(slen),
	}

	currval = false
	for i := 0; i < slen; {
		if s[i] == currval {
			bs.m[idx]++
			i++
			continue
		}
		currval = !currval
		idx++
	}

	return bs
}

// ToSlice converts a BoolSlice to a []bool
func (bs *BoolSlice) ToSlice() []bool {
	bs.rw.RLock()
	defer bs.rw.RUnlock()

	var i, j, cursor int64
	var currval bool

	b := make([]bool, bs.size)

	for ; i < bs.msize; i++ {
		for j = 0; j < bs.m[i]; j++ {
			b[cursor] = currval
			cursor++
		}
		currval = !currval
	}

	return b
}
