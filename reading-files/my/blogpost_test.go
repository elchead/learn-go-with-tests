package main

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestPostFromFolder(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1`
		secondBody = `Title: Post 2
Description: Description 2`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}
	posts, err := NewPostsFromFS(fs)

	got := posts[0]
	want := Post{Title: "Post 1", Description: "Description 1"}
	assert.Equal(t, want, got)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(posts))
}

func TestErrorOpen(t *testing.T) {
	_, err := NewPostsFromFS(StubFailingFS{})
	assert.Error(t, err)
}
