package main

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
	"flag"
)

func main() {
	//var r string = os.Args[1]

	filePointer := flag.String("f", "", "filename")

	flag.Parse()
	if *filePointer == "" {
		log.Fatal("No filename specified, use -f")
	}


	path, err := filepath.Abs(*filePointer)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Stat(*filePointer)
	if err != nil {
		log.Fatal(err)
	}
	b := float64(file.Size())
	kib := b / 1024
	mib := kib / 1024
	gib := mib / 1024
	fmt.Printf("Information about '%s':\n", path)
	fmt.Printf("Size: %.0fB, %fKiB, %fMiB, %.9fGiB\n", b, kib, mib, gib)
	mode := file.Mode()
	if mode&os.ModeDir != 0 {
		fmt.Println("Is a directory file")
	} else {
		fmt.Println("Is not a directory file")
	}
	if mode.IsRegular() != false {
		fmt.Println("Is a regular file")
	} else {
		fmt.Println("Is not a regular file")
	}
	fmt.Println("Has Unix permission bits:", mode)
	if mode&os.ModeAppend != 0 {
		fmt.Println("Is an append file")
	} else {
		fmt.Println("Is not an append file")
	}
	if mode&os.ModeDevice != 0 {
		fmt.Println("Is a device file")
		if mode&os.ModeCharDevice != 0 {
			fmt.Println("Is a Unix character device file")
			fmt.Println("Is not a Unix block device file")
		} else {
			fmt.Println("Is not a Unix character device file")
			fmt.Println("Is a Unix block device file")
		}
	} else {
		fmt.Println("Is not a device file")
		fmt.Println("Is not a Unix character device file")
		fmt.Println("Is not a Unix block device file")
	}
	if mode&os.ModeSymlink != 0 {
		fmt.Println("Is a symbolic link")
	} else {
		fmt.Println("Is not a symbolic link")
	}



}