package main

import "./lineshift"
import "os"

func main() {
	var a = os.Args[1]
	lineshift.FileToByteslice(a)
}
