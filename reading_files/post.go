package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title       string
	Description string
}

// newPost parses the file and returns it
func newPost(postFile io.Reader) (Post, error) {
	// Scanner provides a convenient interface for reading data
	// such as a file of newline-delimited lines of text.
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	title := readLine()[7:]
	description := readLine()[13:]

	return Post{Title: title, Description: description}, nil
}
