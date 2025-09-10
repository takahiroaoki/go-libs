package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPadStart(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "ABCABCAabc", PadStart("abc", "ABC", 10))
}

func TestPadEnd(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "abcABCABCA", PadEnd("abc", "ABC", 10))
}

func TestCalcPadding(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "ABCABCA", calcPadding("abc", "ABC", 10))
}
