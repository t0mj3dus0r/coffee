package collectors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t0mj3dus0r/coffee/streams"
	"github.com/t0mj3dus0r/coffee/streams/collectors"
)

type Person struct {
	Name string
	Age  int
}

func TestGroupingBy(t *testing.T) {
	s := streams.FromArray([]Person{
		{
			Name: "John",
			Age:  57,
		},
		{
			Name: "Dave",
			Age:  57,
		},
		{
			Name: "Alice",
			Age:  45,
		},
		{
			Name: "Berthe",
			Age:  45,
		},
	})

	result := streams.Collect(s,
		collectors.GroupingBy(func(p Person) int { return p.Age }))

	assert.Equal(t, map[int][]Person{
		57: {
			{
				Name: "John",
				Age:  57,
			},
			{
				Name: "Dave",
				Age:  57,
			},
		},
		45: {
			{
				Name: "Alice",
				Age:  45,
			},
			{
				Name: "Berthe",
				Age:  45,
			},
		},
	}, result)
}
