package main

import (
	"embed"
	"flag"
	"fmt"
	"path"

	"github.com/monaco-io/ASCII_ART/font"
)

var (
	//go:embed logo/*
	content embed.FS

	name string
)

func init() {
	flag.StringVar(&name, "name", "bilibili.com", "name of the company logo")
	flag.Parse()
}

func main() {
	logo, err := content.ReadFile(path.Join("logo", name))
	if err != nil {
		logo = []byte(font.AsciiArt(name, ""))
	}
	fmt.Printf("%s", logo)
}
