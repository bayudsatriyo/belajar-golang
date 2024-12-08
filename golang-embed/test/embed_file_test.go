package golangembed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed kobo.png
var kobo []byte

func TestEmbedFile(t *testing.T) {
	err := ioutil.WriteFile("kobo_next.png", kobo, fs.ModePerm)

	if err != nil {
		panic(err)
	}
}
//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")

	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			fmt.Println(dirEntry.Name())
			file, _ := path.ReadFile( "files/" + dirEntry.Name())
			fmt.Println(string(file))
		}
	}

}