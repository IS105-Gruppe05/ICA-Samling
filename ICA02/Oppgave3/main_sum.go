package main

import "fmt"
import "os"
import "strconv"
import "bufio"
import "./sum"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	var text2 string
	fmt.Println("Skriv inn to tall for å summere, skriv q i begge felter for å avslutte")

	for text != "q" {
		fmt.Print("Tall nummer 1: ")
		scanner.Scan()
		text = scanner.Text()
		fmt.Print("Tall nummer 2: ")
		scanner.Scan()
		text2 = scanner.Text()
		if text != "q" {

			tall1, err := strconv.ParseInt(text, 10, 64)
			if err != nil {
				fmt.Println("Tallet kan ikke være: " + text + ", bruk int64 verdi.")
			} else {
				tall2, err := strconv.ParseInt(text2, 10, 64)
				if err != nil {
					fmt.Println("Tallet kan ikke være: " + text2 + ", bruk int64 verdi.")
				} else {
					fmt.Println(text+" + "+text2+" =", sum.SumInt64(tall1, tall2))
				}
			}
		}
	}
}
