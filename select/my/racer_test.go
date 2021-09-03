package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"

	want := fastURL
	got := Racer(slowURL, fastURL)

	assert.Equal(t, want, got)
}
