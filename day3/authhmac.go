package main

import (
	"bufio"
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
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
	fmt.Println(cipher)
	BobDecrypt(cipher, privateKeyB, secret)
}

func AliceEncrypt(msg string, publicKey rsa.PublicKey, secret string) string {
	cipher, _ := rsa.EncryptOAEP(md5.New(), rand.Reader, &publicKey, []byte(msg), nil)
	h := hmac.New(md5.New, []byte(secret))
	h.Write(cipher)
	mac := hex.EncodeToString(h.Sum(nil))
	mac += string(cipher)
	return mac
}

func BobDecrypt(cipher string, privateKey *rsa.PrivateKey, secret string) {
	encrypt := []byte(cipher)
	h := hmac.New(md5.New, []byte(secret))
	h.Write(encrypt[32:])
	mac1 := hex.EncodeToString(h.Sum(nil))
	if string(encrypt[0:32]) == mac1 {
		text, _ := privateKey.Decrypt(nil, encrypt[32:], &rsa.OAEPOptions{Hash: crypto.MD5})
		fmt.Println("Plain Text:", string(text))
	} else {
		fmt.Println(false)
	}
}
