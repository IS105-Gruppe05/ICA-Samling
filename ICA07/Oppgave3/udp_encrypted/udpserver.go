package main

import (
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {

	_, err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("message recieved from %v \n %s ", remoteaddr, p)

		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue

		}
		/*	key := []byte("AES256Key-32Characters1234567890")
			ciphertext, _ := hex.DecodeString("c6805593e5911e3e7f12dd44c4642e647e0f4c17184632c7")

			nonce, _ := hex.DecodeString("7413e381df7cbf94fd01b4dd")

			block, err := aes.NewCipher(key)
			if err != nil {
				panic(err.Error())
			}

			aesgcm, err := cipher.NewGCM(block)
			if err != nil {
				panic(err.Error())
			}

			plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
			if err != nil {
				panic(err.Error())
			}

			fmt.Printf("%s\n", plaintext)
			// Output: exampleplaintext
		*/
		go sendResponse(ser, remoteaddr)
	}

}
