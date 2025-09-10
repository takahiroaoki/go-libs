package slice

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirst(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		elems []string
		f     func(idx int, elem string) bool
		want  int
	}{
		{
			name:  "found",
			elems: []string{"a", "b", "c"},
			f: func(idx int, elem string) bool {
				return elem == "b"
			},
			want: 1,
		},
		{
			name:  "not found",
			elems: []string{"a", "b", "c"},
			f: func(idx int, elem string) bool {
				return elem == "d"
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, FindFirst(tt.elems, tt.f))
		})
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()
	assert.Equal(
		t,
		[]int{-1, -3},
		Filter(
			[]int{1, -1, 2, -3, 0},
			func(idx int, elem int) bool {
				return elem < 0
			},
		),
	)
}

func TestForEach(t *testing.T) {
	t.Parallel()

	bucket := []string{}
	ForEach(
		[]int{1, -1, 2, -3, 0},
		func(idx int, elem int) {
			bucket = append(bucket, strconv.Itoa(elem))
		},
	)
	assert.Equal(
		t,
		[]string{"1", "-1", "2", "-3", "0"},
		bucket,
	)
}

func TestMap(t *testing.T) {
	t.Parallel()

	assert.Equal(
		t,
		[]string{"1", "-1", "2", "-3", "0"},
		Map(
			[]int{1, -1, 2, -3, 0},
			func(idx int, elem int) string {
				return strconv.Itoa(elem)
			},
		),
	)
}
