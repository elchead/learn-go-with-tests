package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	t.Run("increment counter by 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assert.Equal(t, 3, counter.Value())
	})
}
