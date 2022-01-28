package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	key := "Hello"
	text := []byte("This is a message")
	hash := []byte(createHash(key))
	key = string(hash[0:8])
	iv := []byte("abcde125")
	ciphertext := encrypt(text, key, iv)
	textHash := createHash(string(text))
	textHash += string(ciphertext)
	plaintext := decrypt([]byte(textHash), key, iv)
	plainHash := createHash(plaintext)
	myhash := []byte(textHash)
	hash1 := string(myhash[:32])
	if hash1 == plainHash {
		fmt.Println("True")
		fmt.Println("Plain Text:", plaintext)
	} else {
		fmt.Println("False")
	}
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(text []byte, key string, iv []byte) []byte {
	block, _ := des.NewCipher([]byte(key))
	mode := cipher.NewCBCEncrypter(block, iv)
	text = PKCS5Padding(text, block.BlockSize())
	ciphertext := make([]byte, len(text))
	mode.CryptBlocks(ciphertext, text)
	return ciphertext
}

func decrypt(cipherText []byte, key string, iv []byte) string {
	block, _ := des.NewCipher([]byte(key))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(cipherText[32:]))
	decrypter.CryptBlocks(decrypted, cipherText[32:])
	decrypted = PKCS5UnPadding(decrypted)
	return string(decrypted)
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
