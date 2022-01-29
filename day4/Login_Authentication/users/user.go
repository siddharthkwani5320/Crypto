package users

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"login/random"
	"os"
	"strings"
)

type user struct {
	firstname string
	lastname  string
	userid    string
	password  string
	uid       string
}

var allUsers = make(map[string]*user)

func New(fname, lname, id, pass string) {
	var newUser user
	newUser.firstname = fname
	newUser.lastname = lname
	newUser.userid = id
	var uid string
	for {
		uid = random.RandomId()
		_, ok := allUsers[uid]
		if ok {
			continue
		}
		break
	}
	newUser.uid = uid
	newUser.password = createHash(pass + uid)
	allUsers[uid] = &newUser
}

func createHash(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ListAllUser() {
	for _, val := range allUsers {
		fmt.Print("First Name:", val.firstname)
		fmt.Print(" Last Name:", val.lastname)
		fmt.Print(" Userid:", val.userid)
		fmt.Print(" Password:", val.password)
		fmt.Println()
	}
}

func ReadData() {
	var _, err = os.Stat("C:/Users/siddh/OneDrive/Documents/Login_Authentication/data.txt")
	if err != nil {
		f, _ := os.Create("C:/Users/siddh/OneDrive/Documents/Login_Authentication/data.txt")
		f.Close()
	}
	f, err := os.Open("C:/Users/siddh/OneDrive/Documents/Login_Authentication/data.txt")
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
		storeData(data)
	}
}

func storeData(data string) {
	dataSlice := strings.Split(data, ",")
	uid := dataSlice[0]
	var tempUser *user
	tempUser.firstname = dataSlice[1]
	tempUser.lastname = dataSlice[2]
	tempUser.userid = dataSlice[3]
	tempUser.password = dataSlice[4]
	allUsers[uid] = tempUser
}

func WriteData() {
	f, err := os.OpenFile("C:/Users/siddh/OneDrive/Documents/Login_Authentication/data.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var data string
	for key, val := range allUsers {
		data = key + "," + val.firstname + "," + val.lastname + "," + val.userid + "," + val.password + ","
		f.WriteString(data + "\n")

	}
}

func CheckUser(userid, pass string) bool {
	for key, val := range allUsers {
		checkpass := createHash(pass + key)
		if checkpass == val.password {
			return true
		}

	}
	return false
}
