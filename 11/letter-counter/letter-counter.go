package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

// countLetters counts letters from any type that implements the io.Reader
// interface (i.e. has a Read(byte[]) method).
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048) // create buffer once (single alloc)
	out := map[string]int{}
	for {
		// Read() returns number of bytes read.
		n, err := r.Read(buf)
		// iterate over 0:n subslice of buffer
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

// buildGZipReader is a gzip.Reader factory.
func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName) // create *os.File
	if err != nil {
		return nil, nil, err
	}
	// create new *gzip.Reader instance.
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	// return gzip reader and a closure for cleaning up resources when invoked
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func main() {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(counts)

	// now get counts from gzipped file
	r, closer, err := buildGZipReader("my_data.txt.gz")
	if err != nil {
		fmt.Println("error")
	}
	defer closer()
	counts, err = countLetters(r)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(counts)
}
