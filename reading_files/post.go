package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
	bodySeperator        = "Body: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

// newPost parses the file and returns it
func newPost(postBody io.Reader) (Post, error) {
	// Scanner provides a convenient interface for reading data
	// such as a file of newline-delimited lines of text.
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil
}

// readBody read the content body
func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore line
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
