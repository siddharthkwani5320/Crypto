package files

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type file struct {
	name     string
	bell_lvl uint8
	biba_lvl uint8
}

var allFiles []*file

func Read() {
	allFiles = nil
	f, err := os.Open("C:/Users/siddh/OneDrive/Documents/Access Control Model/filedata.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		storeData(strings.TrimSpace(data))
	}
}

func storeData(data string) {
	dataSlice := strings.Split(data, ",")
	var temp file
	temp.name = dataSlice[0]
	level, _ := strconv.Atoi(dataSlice[1])
	temp.bell_lvl = uint8(level)
	level, _ = strconv.Atoi(dataSlice[2])
	temp.biba_lvl = uint8(level)
	allFiles = append(allFiles, &temp)
}

var read []*file
var write []*file

func ReadAble(bell, biba uint8) int {
	read = nil
	for _, val := range allFiles {
		if val.bell_lvl <= bell && val.biba_lvl >= biba {
			fmt.Println(val.name)
			read = append(read, val)
		}
	}
	return len(read)
}

func WriteAble(bell, biba uint8) int {
	write = nil
	for _, val := range allFiles {
		if val.bell_lvl >= bell && val.biba_lvl <= biba {
			fmt.Println(val.name)
			write = append(write, val)
		}
	}
	return len(write)
}

func ReadFile(name string) error {
	for _, val := range read {
		if val.name == name {
			readFile(name)
			return nil
		}
	}
	return errors.New("file not in the given list")
}

func readFile(name string) {
	data, _ := ioutil.ReadFile("files/" + name)
	for _, val := range strings.Split(string(data), "line") {
		if len(val) > 0 {
			plainText := string(decrypt([]byte(strings.TrimSpace(val))))
			fmt.Println(plainText)
		}
	}
}

func WriteFile(name string, phrase string) error {
	for _, val := range write {
		if val.name == name {
			f, err := os.OpenFile("files/"+name, os.O_APPEND, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			cipherText := encrypt([]byte(phrase))
			_, err = f.WriteString("line" + string(cipherText) + "\n")
			if err != nil {
				log.Panic(err)
			}
			return nil
		}
	}
	return errors.New("file not in the given list")
}

func createHash(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte) []byte {
	key := createHash("Siddharth")
	block, _ := aes.NewCipher([]byte(key))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}

func decrypt(data []byte) []byte {
	key := createHash("Siddharth")
	block, _ := aes.NewCipher([]byte(key))
	gcm, _ := cipher.NewGCM(block)
	nonce := data[:gcm.NonceSize()]
	cipherText := data[gcm.NonceSize():]
	plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
	return plainText
}
