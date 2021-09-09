package main

import (
	"bufio"
	"io"
	"io/fs"
)

type Post struct {
	Title       string
	Description string
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
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	title := scanner.Text()
	scanner.Scan()
	description := scanner.Text()
	return Post{Title: string(title[7:]), Description: string(description[14:])}, nil
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
