package main

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	file, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	return Post{Title: string(data[7:])}, err
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, entry := range dir {
		if post, err := getPost(fileSystem, entry); err == nil {
			posts = append(posts, post)
		}

	}
	return posts, err
}
