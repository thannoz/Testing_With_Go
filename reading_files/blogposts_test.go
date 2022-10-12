package blogposts_test

import (
	blogposts "reading_files"
	"reflect"
	"testing"
	"testing/fstest"
)

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello,
world`
	secondBody = `Title: Post 2	
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
)

func TestNewBlogPost(t *testing.T) {

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}
	posts, err := blogposts.NewPostFromFS(fs)
	assertError(t, err)

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
