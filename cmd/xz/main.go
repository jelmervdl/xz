package main

import (
	"io"
	"log"
	"os"
	"github.com/jelmervdl/xz"
)

func main() {
	// create an xz.Reader to decompress the data
	r, err := xz.NewReader(os.Stdin, 0)
	if err != nil {
		log.Fatal(err)
	}
	// write the decompressed data to os.Stdout
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}
