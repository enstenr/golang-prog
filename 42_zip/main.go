package main

import (
	"archive/zip"
 
	"fmt"
	"log"
	"os"
)

func main() {
	archive, err := os.Create("archive.zip")
    if err != nil {
        panic(err)
    }
    defer archive.Close()
    
	// Create a buffer to write our archive to.
 
	// Create a new zip archive.
	w := zip.NewWriter(archive)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := os.Create(file.Name)
		finfo,_:=os.Stat(f.Name())

		fh,_ :=zip.FileInfoHeader(finfo)
		iowriter,_:=w.CreateHeader(fh)
		if err != nil {
			log.Fatal(err)
		}
		_, err = iowriter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(file)
	}
	 

	// Make sure to check the error on Close.
	err 	= w.Close()
	if err != nil {
		log.Fatal(err)
	}
}