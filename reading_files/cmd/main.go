package main

import (
	"fmt"
	"log"
	"os"
	blogposts "reading_files"
)

func main() {
	posts, err := blogposts.NewPostFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(posts)
}
