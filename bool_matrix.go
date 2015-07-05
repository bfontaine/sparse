package sparse

// BoolMatrix is a matrix filled with boolean values.
type BoolMatrix struct {
	bs   *BoolSlice
	w, h int64
}

// NewBoolMatrix returns a pointer on a new BoolMatrix
func NewBoolMatrix(w, h int64) *BoolMatrix {
	m := &BoolMatrix{
		bs: NewBoolSlice(),
		w:  w,
		h:  h,
	}

	return m
}

// Width returns the width of the matrix.
func (m BoolMatrix) Width() int64 { return m.w }

// Height returns the height of the matrix.
func (m BoolMatrix) Height() int64 { return m.h }

// Set sets the value at the given coordinates. This is not implemented right
// now and will return an error.
func (m *BoolMatrix) Set(x, y int64, v bool) error {
	return m.bs.Set(m.index(x, y), v)
}

// Get gets the value at the given coordinates.
func (m *BoolMatrix) Get(x, y int64) bool {
	return m.bs.Get(m.index(x, y))
}

func (m *BoolMatrix) index(x, y int64) int64 {
	return y*m.h + x
}
