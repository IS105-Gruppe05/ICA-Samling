package unicode

import "fmt"

// Kode for Oppgave 4a
const nor = "\x93\x6e\x6f\x72\x64\x20\x6f\x67\x20\x73\xf8\x72"
const is = "\u0022\x6e\x6f\x72\u00f0\x75\x72\x20\x6f\x67\x20\x73\x75\u00f0\x75\x72\u0022"
const jp = "\u0022\xe5\x8c\x97\xE3\x81\xA8\xE5\x8D\x97\u0022"

//language := "“norður og suður”

func Translate(nor string, language string) string {
	if nor == "nord og sør" {
		if language == "jp" {
			fmt.Printf("%s", jp)
		}
		if language == "is" {
			fmt.Printf("%s", is)
		}

	}
	return ""
}
