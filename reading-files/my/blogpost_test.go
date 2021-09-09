package main

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestPostFromFolder(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}
	posts, err := NewPostsFromFS(fs)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(posts))
}
