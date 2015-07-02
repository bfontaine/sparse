package sparse

type BoolSlice struct {
	m     []int64
	msize int64
	size  int64
}

func NewBoolSlice() *BoolSlice {
	return &BoolSlice{
		m:     make([]int64, 1),
		msize: 1,
	}
}

func (bs *BoolSlice) Size() int64 {
	return bs.size
}

func (bs *BoolSlice) Get(idx int64) bool {
	var i, cursor int64

	for ; i < bs.msize; i++ {
		cursor += bs.m[i]
		if idx < cursor {
			break
		}
	}

	return bs.mvalue(i)
}

func (bs *BoolSlice) Set(idx int64, v bool) {
	// TODO
}

func (bs *BoolSlice) Append(v bool) {
	if bs.lastmvalue() == v {
		bs.m[bs.msize-1]++
	} else {
		bs.m = append(bs.m, 1)
		bs.msize++
	}
	bs.size++
}

func (bs *BoolSlice) lastmvalue() bool    { return bs.mvalue(bs.msize - 1) }
func (bs *BoolSlice) mvalue(i int64) bool { return i&1 == 1 }

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

func (bs *BoolSlice) ToSlice() []bool {
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
