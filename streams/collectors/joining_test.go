package collectors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t0mj3dus0r/coffee/streams"
	"github.com/t0mj3dus0r/coffee/streams/collectors"
)

func TestJoining(t *testing.T) {
	s := streams.FromArray([]string{
		"b", "c", "d",
	})

	result := streams.Collect(s, collectors.Joining[string](", ", "a, ", " e"))

	assert.Equal(t, "a, b, c, d e", result)
}
