package main

import "./unicode"
import "os"

//nor := "\x93\x6e\x6f\x72\x64\x20\x6f\x67\x20\x73\xf8\x72"
//is := "\x93\x6e\x6f\xf5\x75\x72\x20\x6f\x20\x73\x75\xf5\x75\x73\x93"
//jp := "\xE3\x81\xA8\xE5\x8D\x97\xe5\x8c\x97"
//language := "“norður og suður”
//Må huske å ha med argument når man kjører programmet. 
func main() {

	var t string = os.Args[1]
	unicode.Translate("nord og sør", t)

}
