package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// guard for file arg
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	// attempt to open file at path - read-only file handle
	f, err := os.Open(os.Args[1])
	// error out if any file opening issues
	if err != nil {
		log.Fatal(err)
	}
	// set up cleanup - close the file regardless of how it's exited!
	defer f.Close()
	// create data byte array that's 2048 bytes long
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}

}
