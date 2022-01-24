package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, phase string) []byte {
	key := createHash(phase)
	block, _ := aes.NewCipher([]byte(key))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}

func decrypt(data []byte, phase string) []byte {
	key := []byte(createHash(phase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonce := data[:gcm.NonceSize()]
	cipherText := data[gcm.NonceSize():]
	plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
	return plainText

}
func main() {
	cipherText := encrypt([]byte("Siddharth"), "hello")
	fmt.Println(string(cipherText))
	plainText := decrypt(cipherText, "hello")
	fmt.Println(string(plainText))
}
