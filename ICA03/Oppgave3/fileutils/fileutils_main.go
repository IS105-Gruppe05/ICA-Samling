package main

import "./fileutils"

func main() {
	fileutils.FileToByteslice("../files/lang01.wl")
	fileutils.FileToByteslice("../files/lang02.wl")
	fileutils.FileToByteslice("../files/lang03.wl")
}
