package sparse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertEqualBoolSlices(t *testing.T, b1, b2 []bool) {
	l := len(b1)
	assert.Equal(t, l, len(b2), "The lengths of %v and %v must be equal", b1, b2)
	for i := 0; i < l; i++ {
		assert.Equal(t, b1[i], b2[i], fmt.Sprintf("The elements %d of %v and %v must be equal", i, b1, b2))
	}
}

func TestBoolSliceFromEmptySlice(t *testing.T) {
	bs := BoolSliceFromSlice([]bool{})
	assert.Equal(t, 1, bs.msize)
	assert.Equal(t, 0, bs.m[0])
	assert.Equal(t, 0, bs.size)
}

func TestBoolSliceFromFalseOnlySlice1(t *testing.T) {
	bs := BoolSliceFromSlice([]bool{false})
	assert.Equal(t, 1, bs.msize)
	assert.Equal(t, 1, bs.m[0])
	assert.Equal(t, 1, bs.size)
}

func TestBoolSliceFromFalseOnlySlice3(t *testing.T) {
	bs := BoolSliceFromSlice([]bool{false, false, false})
	assert.Equal(t, 1, bs.msize)
	assert.Equal(t, 3, bs.m[0])
	assert.Equal(t, 3, bs.size)
}

func TestBoolSliceFromTrueOnlySlice1(t *testing.T) {
	bs := BoolSliceFromSlice([]bool{true})
	assert.Equal(t, 2, bs.msize)
	assert.Equal(t, 0, bs.m[0])
	assert.Equal(t, 1, bs.m[1])
	assert.Equal(t, 1, bs.size)
}

func TestBoolSliceFromTrueOnlySlice3(t *testing.T) {
	bs := BoolSliceFromSlice([]bool{true, true, true})
	assert.Equal(t, 2, bs.msize)
	assert.Equal(t, 0, bs.m[0])
	assert.Equal(t, 3, bs.m[1])
	assert.Equal(t, 3, bs.size)
}

func TestBoolSliceFromToSlice(t *testing.T) {
	for _, b := range [][]bool{
		[]bool{},
		[]bool{true},
		[]bool{false},
		[]bool{true, true, true},
		[]bool{false, false, false},
		[]bool{true, false, false, false},
		[]bool{false, true, true, true},
		[]bool{true, false, true, false, true, false},
		[]bool{false, true, false, true, false, true, false},
	} {
		assertEqualBoolSlices(t, b, BoolSliceFromSlice(b).ToSlice())
	}
}

func TestBoolSliceMValue(t *testing.T) {
	bs := NewBoolSlice()

	assert.Equal(t, false, bs.mvalue(0))
	assert.Equal(t, true, bs.mvalue(1))
	assert.Equal(t, false, bs.mvalue(2))
	assert.Equal(t, true, bs.mvalue(3))
}

func TestBoolSliceMIndex(t *testing.T) {
	assert.Equal(t, 0, BoolSliceFromSlice([]bool{false}).mindex(0))
	assert.Equal(t, 1, BoolSliceFromSlice([]bool{true}).mindex(0))

	assert.Equal(t, 0, BoolSliceFromSlice([]bool{false, false, true}).mindex(0))
	assert.Equal(t, 0, BoolSliceFromSlice([]bool{false, false, true}).mindex(1))
	assert.Equal(t, 1, BoolSliceFromSlice([]bool{false, false, true}).mindex(2))
}

func TestBoolSliceGet(t *testing.T) {
	bs := BoolSliceFromSlice([]bool{false, true, true, false, false, true, false})

	assert.Equal(t, false, bs.Get(0))
	assert.Equal(t, true, bs.Get(1))
	assert.Equal(t, true, bs.Get(2))
	assert.Equal(t, false, bs.Get(3))
}

func TestBoolSliceAppendStartWithTrue(t *testing.T) {
	bs := NewBoolSlice()
	assert.Equal(t, 0, bs.Size())

	bs.Append(true)
	assert.Equal(t, 1, bs.Size())
	assert.Equal(t, true, bs.Get(0))

	bs.Append(true)
	assert.Equal(t, 2, bs.Size())
	assert.Equal(t, true, bs.Get(1))

	bs.Append(false)
	assert.Equal(t, 3, bs.Size())
	assert.Equal(t, false, bs.Get(2))
}

func TestBoolSliceAppendStartWithFalse(t *testing.T) {
	bs := NewBoolSlice()
	assert.Equal(t, 0, bs.Size())

	bs.Append(false)
	assert.Equal(t, 1, bs.Size())
	assert.Equal(t, false, bs.Get(0))

	bs.Append(false)
	assert.Equal(t, 2, bs.Size())
	assert.Equal(t, false, bs.Get(1))

	bs.Append(true)
	assert.Equal(t, 3, bs.Size())
	assert.Equal(t, true, bs.Get(2))
}

func TestBoolSliceSetNoop(t *testing.T) {
	bs := BoolSliceFromSlice([]bool{true})
	bs.Set(0, true)
	assert.Equal(t, true, bs.Get(0))
}

func TestBoolSliceSimpleSet(t *testing.T) {
	bs := BoolSliceFromSlice([]bool{true})
	bs.Set(0, false)
	assert.Equal(t, false, bs.Get(0))
	bs.Set(0, true)
	assert.Equal(t, true, bs.Get(0))

	bs.Append(true)
	assert.Equal(t, true, bs.Get(0))
	assert.Equal(t, true, bs.Get(1))

	bs.Set(1, false)
	assert.Equal(t, true, bs.Get(0))
	assert.Equal(t, false, bs.Get(1))

	bs.Set(0, false)
	assert.Equal(t, false, bs.Get(0))
	assert.Equal(t, false, bs.Get(1))

	bs.Set(1, true)
	assert.Equal(t, false, bs.Get(0))
	assert.Equal(t, true, bs.Get(1))
}

func TestBoolSliceSet1(t *testing.T) {
	bs := BoolSliceFromSlice([]bool{false, false, false, true, false, false})
	bs.Set(1, true)

	assertEqualBoolSlices(t, []bool{false, true, false, true, false, false}, bs.ToSlice())
}

func TestBoolSliceSet2(t *testing.T) {
	bs := BoolSliceFromSlice([]bool{true, true, false, false})
	bs.Set(1, false)

	assertEqualBoolSlices(t, []bool{true, false, false, false}, bs.ToSlice())
}

func TestBoolSliceSet3(t *testing.T) {
	bs := BoolSliceFromSlice([]bool{true, true, false, false})
	bs.Set(1, false)

	assertEqualBoolSlices(t, []bool{true, false, false, false}, bs.ToSlice())
}
