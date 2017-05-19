package lineshift

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func FileToByteslice(filename string) []byte {

	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	lineBreak2 := ([]byte("\x0A"))
	lineBreak1 := ([]byte("\x0D\x0A"))

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("% X %c", byteSlice, byteSlice)
	fmt.Println()
	if bytes.Contains(byteSlice, lineBreak1) {
		fmt.Println("Det er", bytes.Count(byteSlice, lineBreak1), "linjeskift i denne filen.(Carriage return og lineshift)")
	} else if bytes.Contains(byteSlice, lineBreak2) {
		fmt.Println("Det er", bytes.Count(byteSlice, lineBreak2), "linjeskift i denne filen.(Lineshift)")
	} else {
		fmt.Println("Det er ingen linjeskift i filen.")
	}
	return byteSlice
}
