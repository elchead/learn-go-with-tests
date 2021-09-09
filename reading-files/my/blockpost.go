package main

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

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
	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagsSeparator), ", ")
	return Post{Title: title, Description: description, Tags: tags}, nil
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
