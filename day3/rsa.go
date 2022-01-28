package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
)

func main() {
	privateKeyB, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKeyB := privateKeyB.PublicKey

	msg := "Siddharth"
	cipher, _ := rsa.EncryptOAEP(sha512.New(), rand.Reader, &publicKeyB, []byte(msg), nil)

	fmt.Println("Cipher Text:", string(cipher))

	text, _ := privateKeyB.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA512})
	fmt.Println("Plain Text:", string(text))
}
