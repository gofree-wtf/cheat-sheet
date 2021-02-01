package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Card struct {
	Cost int
}

func (c *Card) SubCost(n int) {
	c.Cost -= n
}

func TestForEach1(t *testing.T) {
	deck := []*Card{
		{Cost: 10},
		{Cost: 5},
	}

	for _, c := range deck {
		c.SubCost(5)
	}

	assert.Equal(t, 5, deck[0].Cost)
	assert.Equal(t, 0, deck[1].Cost)
}

func TestForEach2(t *testing.T) {
	deck := []Card{
		{Cost: 10},
		{Cost: 5},
	}

	for i := range deck {
		c := &deck[i]
		c.SubCost(5)
	}

	assert.Equal(t, 5, deck[0].Cost)
	assert.Equal(t, 0, deck[1].Cost)
}
