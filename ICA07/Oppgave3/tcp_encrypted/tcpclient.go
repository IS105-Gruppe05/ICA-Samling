package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	localAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	// connect to 127.0.0.1:9090
	conn, err := net.DialTCP("tcp", nil, localAddr)
	if err != nil {
		errors.New("Problem connecting.")
	}
	// read from stdin
	reader := bufio.NewReader(os.Stdin)
	key := []byte("AES256Key-32Characters1234567890")
	for {

		fmt.Print("> ")
		msg, _ := reader.ReadBytes(0xA)
		// Kill the newline char
		msg = msg[:len(msg)-1]

		block, err := aes.NewCipher(key)
		if err != nil {
			panic(err.Error())
		}

		nonce := make([]byte, 12)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			panic(err.Error())
		}

		aesgcm, err := cipher.NewGCM(block)
		if err != nil {
			panic(err.Error())
		}

		ciphertext := aesgcm.Seal(nil, nonce, msg, nil)

		// write to the connection
		conn.Write(ciphertext)

		response := make([]byte, 1024)
		// read the response
		_, err = conn.Read(response)
		if err != nil {
			fmt.Print("Connection to the server was closed.\n")
			break
		}
		fmt.Printf("%s\n", response)

	}
}
