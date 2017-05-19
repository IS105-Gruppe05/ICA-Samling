package main

import (
	"./jsonurl"
	"./jsonurlv2"
	"./jsonurlv3"
	"./jsonurlv4"
	"./jsonurlv5"
	"fmt"
	"time"
)

func main() {

	 channel := make(chan bool, 1)
	 go kilde(channel)
	 <-channel
}

func kilde(channel chan bool) {

	fmt.Println("Henter kilder...\n")
	time.Sleep(time.Second)
	jsonurl.JsonUrl()
	time.Sleep(time.Second)
	jsonurlv2.JsonUrlV2()
	time.Sleep(time.Second)
	jsonurlv3.JsonUrlV3()
	time.Sleep(time.Second)
	jsonurlv4.JsonUrlV4()
	time.Sleep(time.Second)
	jsonurlv5.JsonUrlV5()
	fmt.Println("Alle kildene er hentet")

	channel<-true

}
