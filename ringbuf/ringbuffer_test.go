package ringbuf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingbufferPushStartPopStart(t *testing.T) {
	size := 4
	r := New(size)

	// Push values
	for i := 0; i < size; i++ {
		r.PushStart(i)
	}

	// Check can't push when full
	assert.Error(t, r.PushStart(0))

	// Check retrieved values
	for i := size - 1; i >= 0; i-- {
		v, err := r.PopStart()
		assert.NoError(t, err)
		assert.Equal(t, i, v)
	}

	// Check can't pop when empty
	_, err := r.PopStart()
	assert.Error(t, err)
}

func TestRingbufferPushEndPopEnd(t *testing.T) {
	size := 4
	r := New(size)

	// Push values
	for i := 0; i < size; i++ {
		r.PushEnd(i)
	}

	// Check can't push when full
	assert.Error(t, r.PushEnd(0))

	// Check retrieved values
	for i := size - 1; i >= 0; i-- {
		v, err := r.PopEnd()
		assert.NoError(t, err)
		assert.Equal(t, i, v)
	}

	// Check can't pop when empty
	_, err := r.PopEnd()
	assert.Error(t, err)
}

func TestRingbufferPushStartPopEnd(t *testing.T) {
	size := 4
	r := New(size)

	// Push values
	for i := 0; i < size; i++ {
		r.PushStart(i)
	}

	// Check can't push when full
	assert.Error(t, r.PushStart(0))

	// Check retrieved values
	for i := 0; i < size; i++ {
		v, err := r.PopEnd()
		assert.NoError(t, err)
		assert.Equal(t, i, v)
	}

	// Check can't pop when empty
	_, err := r.PopEnd()
	assert.Error(t, err)
}

// func TestRingbufferPushPopRandom(t *testing.T) {
// 	size := 4
// 	r := New(size)

// 	// Push values
// 	for i := 0; i < size; i++ {
// 		r.Push(i, true)
// 	}

// 	// Check can't push when full
// 	assert.Error(t, r.Push(0, true))

// 	// Check retrieved values
// 	for i := 0; i < size; i++ {
// 		v, err := r.PopEnd()
// 		assert.NoError(t, err)
// 		assert.Equal(t, i, v)
// 	}

// 	// Check can't pop when empty
// 	_, err := r.PopEnd()
// 	assert.Error(t, err)
// }
