package main

import "os"
import "./finnbokstav"
import "./lineshift"

func main() {
	var na = os.Args[1]
	oppg3.FinnAlle(na)
	lineshift.FileToByteslice(na)

}
