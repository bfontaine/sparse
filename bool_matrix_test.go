package sparse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolMatrixInitFalse(t *testing.T) {
	m := NewBoolMatrix(2, 4)

	for x := int64(0); x < 2; x++ {
		for y := int64(0); y < 4; y++ {
			assert.Equal(t, false, m.Get(x, y))
		}
	}
}
