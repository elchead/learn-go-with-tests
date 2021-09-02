package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			fn := func(s string) { got = append(got, s) }
			walk(test.Input, fn)
			assert.Equal(t, test.ExpectedCalls, got)
		})
	}
}
