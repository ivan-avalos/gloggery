package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// TODO: Decide paths with command-line options
	homeFolder := os.Getenv("HOME")
	postsFolder := filepath.Join(homeFolder, "gloggery/posts")
	templatesFolder := filepath.Join(homeFolder, "gloggery/templates")
	glogFolder := filepath.Join(homeFolder, "public_gemini/glog")

	filenames := listFolderItemsReverse(postsFolder)

	if len(filenames) == 0 {
		fmt.Printf("no posts in %v\n", postsFolder)
		os.Exit(0)
	}

	templates := parseTemplates(templatesFolder)

	posts := make([]*Post, 0, len(filenames))
	for _, filename := range filenames {
		posts = append(posts, readPost(postsFolder, filename))
	}

	writeGlog(templates, glogFolder, posts)
}
