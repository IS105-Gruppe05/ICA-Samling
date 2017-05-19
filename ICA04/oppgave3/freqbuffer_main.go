package main

import "os"
import "./freqbuffer"
import "./lineshift"

func main() {
	var na = os.Args[1]
	oppg3.ReadBuffer(na)
	lineshift.FileToByteslice(na)
}
