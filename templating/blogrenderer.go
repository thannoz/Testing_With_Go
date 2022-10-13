package blogrenderer

import (
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w *io.Writer, p Post) error {
	return nil
}

func (p *Post) blogrenderer(input io.Writer) {

}
