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
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}
	posts, err := NewPostsFromFS(fs)

	got := posts[0]
	want := Post{Title: "Post 1"}
	assert.Equal(t, want, got)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(posts))
}

func TestErrorOpen(t *testing.T) {
	_, err := NewPostsFromFS(StubFailingFS{})
	assert.Error(t, err)
}
