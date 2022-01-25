package main

import (
	"bufio"
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CreateKey(key string) string {
	byteKey := []byte(createHash(key))
	return string(byteKey[0:8])
}

func encrypt(text []byte, key string, iv []byte) []byte {
	block, _ := des.NewCipher([]byte(key))
	mode := cipher.NewCBCEncrypter(block, iv)
	text = PKCS5Padding(text, block.BlockSize())
	ciphertext := make([]byte, len(text))
	mode.CryptBlocks(ciphertext, text)
	return ciphertext
}

func decrypt(cipherText []byte, key string, iv []byte) []byte {
	block, _ := des.NewCipher([]byte(key))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(cipherText))
	decrypter.CryptBlocks(decrypted, cipherText)
	decrypted = PKCS5UnPadding(decrypted)
	return decrypted
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

func main() {
	var s string
	fmt.Println("Enter your Message:")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	s = strings.TrimSpace(s)
	key, _ := reader.ReadString('\n')
	key = strings.TrimSuffix(key, "\n")
	key = CreateKey(key)
	plaintext := []byte(s)
	fmt.Println(len(plaintext))
	iv := []byte("abcde125")
	cipherText := encrypt(plaintext, key, iv)
	fmt.Printf("%s encrypt to %x \n", plaintext, cipherText)
	decryptedText := decrypt(cipherText, key, iv)
	fmt.Printf("%x decrypt to %s\n", cipherText, decryptedText)
}
