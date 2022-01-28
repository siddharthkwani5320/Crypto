package main

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	secret := "mykey"
	data := "Siddharth"
	h := hmac.New(md5.New, []byte(data))
	h.Write([]byte(secret))
	mac := hex.EncodeToString(h.Sum(nil))
	fmt.Println(mac)
}
