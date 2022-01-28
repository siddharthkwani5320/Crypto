package main

import (
	"bufio"
	"crypto"
	"crypto/aes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/aead/cmac"
)

func main() {
	var s string
	fmt.Println("Enter your Message:")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println("Enter your Secret Key:")
	secret, _ := reader.ReadString('\n')
	secret = strings.TrimSuffix(secret, "\n")
	privateKeyB, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKeyB := privateKeyB.PublicKey
	cipher := AliceEncrypt(s, publicKeyB, secret)
	fmt.Println("Cipher Text:", cipher)
	BobDecrypt(cipher, privateKeyB, secret)
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func AliceEncrypt(msg string, publicKey rsa.PublicKey, secret string) string {
	cipher, _ := rsa.EncryptOAEP(md5.New(), rand.Reader, &publicKey, []byte(msg), nil)
	key := []byte(createHash(secret))
	block, _ := aes.NewCipher(key)
	h, _ := cmac.New(block)
	h.Write(cipher)
	mac := hex.EncodeToString(h.Sum(nil))
	mac += string(cipher)
	return mac
}

func BobDecrypt(cipher string, privateKey *rsa.PrivateKey, secret string) {
	encrypt := []byte(cipher)
	text, _ := privateKey.Decrypt(nil, encrypt[32:], &rsa.OAEPOptions{Hash: crypto.MD5})
	fmt.Println(string(text))
	key := []byte(createHash(secret))
	block, _ := aes.NewCipher(key)
	h, _ := cmac.New(block)
	h.Write(encrypt[32:])
	mac1 := hex.EncodeToString(h.Sum(nil))
	fmt.Println(mac1)
	if string(encrypt[0:32]) == mac1 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
