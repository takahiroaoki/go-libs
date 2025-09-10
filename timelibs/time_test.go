package timelibs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetLocation(t *testing.T) {
	// JST
	SetLocation(time.FixedZone("JST", 9*60*60))
	assert.Equal(t, time.FixedZone("JST", 9*60*60), location)

	ti := NewTime(time.Date(2025, 9, 10, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, time.Date(2025, 9, 10, 9, 0, 0, 0, time.FixedZone("JST", 9*60*60)), ti.Time())

	now := Now()
	assert.Equal(t, time.FixedZone("JST", 9*60*60), now.Time().Location())

	// UTC
	SetLocation(time.UTC)
	assert.Equal(t, time.UTC, location)

	ti = NewTime(time.Date(2025, 9, 10, 0, 0, 0, 0, time.FixedZone("JST", 9*60*60)))
	assert.Equal(t, time.Date(2025, 9, 9, 15, 0, 0, 0, time.UTC), ti.Time())

	now = Now()
	assert.Equal(t, time.UTC, now.Time().Location())
}
