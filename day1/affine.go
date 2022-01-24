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
	fmt.Println("Enter first key:")
	keyTemp1, _ := reader.ReadString('\n')
	keyTemp1 = strings.TrimSpace(keyTemp1)
	key1, _ := strconv.ParseInt(keyTemp1, 10, 64)
	fmt.Println("Enter second key:")
	keyTemp2, _ := reader.ReadString('\n')
	keyTemp2 = strings.TrimSpace(keyTemp2)
	key2, _ := strconv.ParseInt(keyTemp2, 10, 64)
	a := int(key1)
	b := int(key2)
	s = encrypt(s, a, b)
	fmt.Println("Encrypt message:", s)
	s = decrypt(s, a, b)
	fmt.Println("Decrypt message:", s)
}

func encrypt(s string, a int, b int) string {
	var temp string
	for i := 0; i < len(s)-1; i++ {
		temp1 := ((int(s[i]-97) * a) + b) % 26
		char1 := string(temp1 + 97)
		temp += char1
	}
	return temp
}

func decrypt(s string, a int, b int) string {
	var temp string
	inverseA := InverseKey(a)
	for i := 0; i < len(s); i++ {
		var temp1 int
		temp1 = (((int(s[i]-97) * inverseA) % 26) - (b % 26) + 26) % 26
		char1 := string(temp1 + 97)
		temp += char1
	}
	return temp
}

func InverseKey(key int) int {
	var inverseKey int
	for i := 0; i < 26; i++ {
		if ((key * i) % 26) == 1 {
			inverseKey = i
			break
		}
	}
	return inverseKey
}
