package goption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNothing(t *testing.T) {
	assert.True(t, true)
}

func TestNone(t *testing.T) {
	none := NewFloat64None()
	assert.Equal(t, false, none.IsSet())
	assert.Equal(t, -1.0, none.GetOrElse(-1.0))
	_, isSet := none.Get()
	assert.Equal(t, false, isSet)

	assert.Equal(t, []float64{}, none.ToSlice())
}

func TestSome(t *testing.T) {
	obj := NewFloat64Some(2.0)
	assert.Equal(t, true, obj.IsSet())
	assert.Equal(t, 2.0, obj.GetOrElse(-1.0))
	value, isSet := obj.Get()
	assert.Equal(t, true, isSet)
	assert.Equal(t, 2.0, value)

	assert.Equal(t, []float64{2.0}, obj.ToSlice())
}

func TestSetReset(t *testing.T) {
	obj := NewFloat64None()
	assert.Equal(t, false, obj.IsSet())
	assert.Equal(t, -1.0, obj.GetOrElse(-1.0))

	obj.Set(1.0)
	assert.Equal(t, true, obj.IsSet())
	assert.Equal(t, 1.0, obj.GetOrElse(-1.0))

	obj.Reset()
	assert.Equal(t, false, obj.IsSet())
}
