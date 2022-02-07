package user

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"model/random"
	"os"
	"strconv"
	"strings"
)

type users struct {
	firstname string
	lastname  string
	userid    string
	password  string
	uid       string
	bell_lvl  uint8
	biba_lvl  uint8
}

var allUsers = make(map[string]*users)

func New(fname, lname, id, pass string, bell_l uint8, biba_l uint8) {
	var newUser users
	newUser.firstname = fname
	newUser.lastname = lname
	newUser.userid = id
	newUser.bell_lvl = bell_l
	newUser.biba_lvl = biba_l
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

func ListAllUser() {
	for _, val := range allUsers {
		fmt.Print("First Name:", val.firstname)
		fmt.Print(" Last Name:", val.lastname)
		fmt.Print(" Bell la Padula Level:", val.bell_lvl)
		fmt.Print(" BIBA Level:", val.biba_lvl)
		fmt.Print(" Userid:", val.userid)
		fmt.Print(" Password:", val.password)
		fmt.Println()
	}
}

func ReadData() {
	var _, err = os.Stat("C:/Users/siddh/OneDrive/Documents/Access Control Model/data.txt")
	if err != nil {
		f, _ := os.Create("C:/Users/siddh/OneDrive/Documents/Access Control Model/data.txt")
		f.Close()
	}
	f, err := os.Open("C:/Users/siddh/OneDrive/Documents/Access Control Model/data.txt")
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
	uid := dataSlice[0]
	var tempUser users
	tempUser.firstname = dataSlice[1]
	tempUser.lastname = dataSlice[2]
	tempUser.userid = dataSlice[3]
	tempUser.password = dataSlice[4]
	level, _ := strconv.Atoi(dataSlice[5])
	tempUser.bell_lvl = uint8(level)
	level, _ = strconv.Atoi(dataSlice[6])
	tempUser.biba_lvl = uint8(level)
	allUsers[uid] = &tempUser
}

func WriteData() {
	f, err := os.OpenFile("C:/Users/siddh/OneDrive/Documents/Access Control Model/data.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var data string
	for key, val := range allUsers {
		data = key + "," + val.firstname + "," + val.lastname + "," + val.userid + "," + val.password + "," + fmt.Sprintf("%d", (int(val.bell_lvl))) + "," + fmt.Sprintf("%d", (int(val.biba_lvl))) + ","
		f.WriteString(data + "\n")
	}
}

func CheckUser(userid, pass string) (bool, uint8, uint8) {
	for key, val := range allUsers {
		checkpass := createHash(pass + key)
		if checkpass == val.password && val.userid == userid {
			return true, val.bell_lvl, val.biba_lvl
		}
	}
	return false, 0, 0
}

func createHash(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	return hex.EncodeToString(hasher.Sum(nil))
}
