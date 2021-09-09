package main

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

func getPost(fileSystem fs.FS, filename string) (Post, error) {
	file, err := fileSystem.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()
	return newPost(file)
}

func newPost(f io.Reader) (Post, error) {
	data, err := io.ReadAll(f)
	return Post{Title: string(data[7:])}, err
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, entry := range dir {
		if post, err := getPost(fileSystem, entry.Name()); err == nil {
			posts = append(posts, post)
		}

	}
	return posts, err
}
