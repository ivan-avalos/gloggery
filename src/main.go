package main

import (
	"os"
	"path/filepath"
)

func main() {
	// TODO: Decide paths with command-line options
	homeFolder := os.Getenv("HOME")
	postsFolder := filepath.Join(homeFolder, "gloggery/posts")
	templatesFolder := filepath.Join(homeFolder, "gloggery/templates")
	glogFolder := filepath.Join(homeFolder, "public_gemini/glog")

	title := "~easeout"
	glogURL := "gemini://tilde.team/~easeout/glog"

	postIndex := make(chan *postIndex)
	go loadPostIndex(postsFolder, title, glogURL, postIndex)

	builder := make(chan *builder)
	go loadBuilder(templatesFolder, builder)

	(<-builder).buildGlog(glogFolder, <-postIndex)
}
