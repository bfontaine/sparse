package sparse

type BoolMatrix struct {
	bs   *BoolSlice
	w, h int64
}

func NewBoolMatrix(w, h int64) *BoolMatrix {
	m := &BoolMatrix{
		bs: NewBoolSlice(),
		w:  w,
		h:  h,
	}

	return m
}

func (m *BoolMatrix) Set(x, y int64, v bool) {
	m.bs.Set(m.index(x, y), v)
}

func (m *BoolMatrix) Get(x, y int64) bool {
	return m.bs.Get(m.index(x, y))
}

func (m *BoolMatrix) index(x, y int64) int64 {
	return y*m.h + x
}
