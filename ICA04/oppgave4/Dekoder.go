package main

import "fmt"
import (
	"strings"
	"bufio"
)




var Helse = 0
var Pedag = 0
var Kunst = 0
var Tek = 0
var Lærer = 0
var Økonomi = 0

func main() {
	in := "001001001001001001001001001001001001001001001001001" +
		"00000000000000000000000000000000000000000011111111111110101010101010101" +
		"01010101010101010101010101101101101101101101101101101101101101101100101" +
		"01010101010101010101010101010101010101010101010101010101"

	s := bufio.NewScanner(strings.NewReader(in))
	s.Split(bufio.ScanRunes)


	for s.Scan() {
		if s.Text() == "0" {
			s.Scan()
			if s.Text() =="0"{
				s.Scan()
				if s.Text() == "1" {
					//fmt.Println("001")
					Helse++
				}else if s.Text() == "0"{
					//fmt.Println("000")
					Pedag++
				}
			}else if s.Text() == "1" {
				//fmt.Println("01")
				Økonomi++
			}


		}else if s.Text() == "1"{
			s.Scan()
			if s.Text() == "1" {
				s.Scan()
				if s.Text() == "1"{
					//fmt.Println("111")
					Kunst++
				}else if s.Text() == "0"{
					//fmt.Println("110")
					Lærer++
				}
			}else if s.Text() == "0"{
				//fmt.Println("10")
				Tek++
			}
		}
	}
	fmt.Println("Helse og idrett: " , Helse, "%")
	fmt.Println("Humaniora og pedagogikk: " , Pedag, "%")
	fmt.Println("Kunstfag: " , Kunst, "%")
	fmt.Println("Teknologi og realfag: " , Tek, "%")
	fmt.Println("Lærerutdanning: " , Lærer, "%")
	fmt.Println("Økonomi og samfunnsvitenskap: " , Økonomi, "%")
}
