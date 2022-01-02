package main

import (
	"fmt"
	"log"
	"os"
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
	"io"
	"errors"
)

func encrypt(key []byte, bytes []byte, fileName string) {
	block, err := aes.NewCipher(key)

	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        fmt.Println(err)
    }

	err = os.WriteFile(fileName, gcm.Seal(nonce, nonce, bytes, nil), 0777)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encryption complete!")

}

func decrypt(key []byte, bytes []byte, fileName string) {
	block, err := aes.NewCipher(key)

    if err != nil {
        fmt.Println(err)
    }

    gcm, err := cipher.NewGCM(block)

    if err != nil {
        fmt.Println(err)
    }

    nonceSize := gcm.NonceSize()
    if len(bytes) < nonceSize {
        fmt.Println(err)
    }

	nonce, bytes := bytes[:nonceSize], bytes[nonceSize:]

	text, err := gcm.Open(nil, nonce, bytes, nil)

	if err != nil {
		log.Fatal(err)
	}
	
	err = os.WriteFile(fileName, text, 0777)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decryption complete!")

}

func main() {

	method := os.Args[1]
	fileName := os.Args[2] 

	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	bytes, err := os.ReadFile(fileName)
	KEY := []byte("hastobe16bytes:)")

	if err != nil {
		log.Fatal(err)

	}

	switch method {
	case "-e":
		encrypt(KEY, bytes, fileName)
	case "-d":
		decrypt(KEY, bytes, fileName)
	}
	
}