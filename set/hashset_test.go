package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHashSetIsEmpty(t *testing.T) {
	set := NewHashSet[string]()

	assert.True(t, set.IsEmpty())
}

func TestAddOne(t *testing.T) {
	set := NewHashSet[string]()

	inserted := set.Add("test")

	assert.True(t, inserted)
}

func TestAddTwiceTheSameValue(t *testing.T) {
	set := NewHashSet[string]()

	inserted := set.Add("test")

	assert.True(t, inserted)

	inserted = set.Add("test")

	assert.False(t, inserted)
}

func TestAddOneThenSizeIsOne(t *testing.T) {
	set := NewHashSet[string]()

	set.Add("test")

	assert.Equal(t, 1, set.Size())
}

func TestAddTwoThenSizeIsTwo(t *testing.T) {
	set := NewHashSet[string]()

	set.Add("test")
	set.Add("test2")

	assert.Equal(t, 2, set.Size())
}

func TestAddOneAndClearThenIsEmpty(t *testing.T) {
	set := NewHashSet[string]()

	set.Add("test")

	set.Clear()

	assert.True(t, set.IsEmpty())
}

func TestToArray(t *testing.T) {
	set := NewHashSet[string]()

	set.Add("test")
	set.Add("test2")
	set.Add("test2")

	assert.Equal(t, []string{"test", "test2"}, set.ToArray())
}
