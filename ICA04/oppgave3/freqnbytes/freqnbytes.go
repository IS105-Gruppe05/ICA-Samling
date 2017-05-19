package oppg3

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const Items = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz.,!?`

func FindNBytes(filename string) {

	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	byteSlice := make([]byte, 4096)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	var biggest int = 0
	var andre int = 0
	var tredje int = 0
	var fjerde int = 0
	var femte int = 0

	var bokstav byte
	var andreb byte
	var tredjeb byte
	var fjerdeb byte
	var femteb byte

	for i := 0; i < len(Items); i++ {

		if bytes.Contains(byteSlice, []byte{Items[i]}) {
			fmt.Printf("Det er %d %q i denne filen\n", bytes.Count(byteSlice, []byte{Items[i]}), Items[i])
			v, _ := bytes.Count(byteSlice, []byte{Items[i]}), Items[i]
			if v > femte {
				femte = v
				femteb = Items[i]
				if fjerde < femte {
					fjerde = femte
					fjerdeb = femteb
					femte = 0

					if tredje < fjerde {
						tredje = fjerde
						tredjeb = fjerdeb
						fjerde = 0
						if andre < tredje {
							andre = tredje
							andreb = fjerdeb
							tredje = 0
							if biggest < andre {
								biggest = andre
								bokstav = andreb
								andre = 0

							}
						}
					}
				}
			}
		} else {

			fmt.Printf("Der er ingen %q i denne filen.\n", Items[i])
		}
	}
	fmt.Println("")
	fmt.Printf("Det er flest av %q i denne filen og det er %d stk\n", bokstav, biggest)
	fmt.Printf("Det er nest flest av %q i denne filen og det er %d stk\n", andreb, andre)
	fmt.Printf("Det er tredje flest av %q i denne filen og det er %d stk\n", tredjeb, tredje)
	fmt.Printf("Det er fjerde flest av %q i denne filen og det er %d stk\n", fjerdeb, fjerde)
	fmt.Printf("Det er femte flest av %q i denne filen og det er %d stk\n", femteb, femte)
	return
}

func FindNBytesTest(filename string) {

	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	byteSlice := make([]byte, 4096)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	var biggest int = 0
	var andre int = 0
	var tredje int = 0
	var fjerde int = 0
	var femte int = 0

	var bokstav byte
	var andreb byte
	var tredjeb byte
	var fjerdeb byte
	var femteb byte

	if err != nil {
		fmt.Println("")
		fmt.Printf("Det er flest av %q i denne filen og det er %d stk\n", bokstav, biggest)
		fmt.Printf("Det er nest flest av %q i denne filen og det er %d stk\n", andreb, andre)
		fmt.Printf("Det er tredje flest av %q i denne filen og det er %d stk\n", tredjeb, tredje)
		fmt.Printf("Det er fjerde flest av %q i denne filen og det er %d stk\n", fjerdeb, fjerde)
		fmt.Printf("Det er femte flest av %q i denne filen og det er %d stk\n", femteb, femte)
		log.Printf("Number of bytes read: %d\n", bytesRead)
		log.Printf("Data read: %s\n", byteSlice)
		return
	}

	for i := 0; i < len(Items); i++ {

		if bytes.Contains(byteSlice, []byte{Items[i]}) {

			v, _ := bytes.Count(byteSlice, []byte{Items[i]}), Items[i]
			if v > femte {
				femte = v
				femteb = Items[i]
				if fjerde < femte {
					fjerde = femte
					fjerdeb = femteb
					femte = 0

					if tredje < fjerde {
						tredje = fjerde
						tredjeb = fjerdeb
						fjerde = 0
						if andre < tredje {
							andre = tredje
							andreb = fjerdeb
							tredje = 0
							if biggest < andre {
								biggest = andre
								bokstav = andreb
								andre = 0
							}
						}
					}
				}
			}
		}
		file.Close()
	}
	return
}
