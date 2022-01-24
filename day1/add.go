package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Println("Enter your Message:")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	data := strings.Split(s, " ")
	s = ""
	for _, i := range data {
		s += i
	}
	fmt.Println("Enter a key:")
	keyTemp, _ := reader.ReadString('\n')
	keyTemp = strings.TrimSpace(keyTemp)
	key1, _ := strconv.ParseInt(keyTemp, 10, 64)
	key := int(key1)
	fmt.Println(key)
	s = encrypt(s, key)
	fmt.Println("Encrypt message:", s)
	s = decrypt(s, key)
	fmt.Println("Decrypt message:", s)
}

func encrypt(s string, key int) string {
	var temp string
	for i := 0; i < len(s)-1; i++ {
		temp1 := (int(s[i]-97) + key) % 26
		char1 := string(temp1 + 97)
		temp += char1
	}
	return temp
}

func decrypt(s string, key int) string {
	var temp string
	for i := 0; i < len(s); i++ {
		var temp1 int
		temp1 = (int(s[i]-97) - (key % 26) + 26) % 26
		char1 := string(temp1 + 97)
		temp += char1
	}
	return temp
}
