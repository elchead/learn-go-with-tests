package main

import (
	"sync"
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

	t.Run("increment concurrently", func(t *testing.T) {
		counter := Counter{}
		wantedCounter := 1000

		var wg sync.WaitGroup
		wg.Add(wantedCounter)
		for i := 0; i < wantedCounter; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assert.Equal(t, wantedCounter, counter.Value())

	})
}
