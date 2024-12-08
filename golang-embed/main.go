package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed kobo.png
var kobo []byte

//go:embed files/*
var path embed.FS

func main()  {
	fmt.Println(version)

	err := ioutil.WriteFile("kobo_next_2.png", kobo, fs.ModePerm)

	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("files")

	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			fmt.Println(dirEntry.Name())
			file, _ := path.ReadFile( "files/" + dirEntry.Name())
			fmt.Println(string(file))
		}
	}
}