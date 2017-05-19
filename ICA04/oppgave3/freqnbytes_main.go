package main

import "os"
import "./freqnbytes"
import "./lineshift"

func main() {
	var na = os.Args[1]
	oppg3.FindNBytes(na)
	lineshift.FileToByteslice(na)

}
