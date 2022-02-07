package main

import (
	"fmt"
	"model/files"
	"model/user"
	"os"
)

func main() {
	user.ReadData()
	for {
		fmt.Println("\tMenu", "\nEnter 1 to Add User", "\nEnter 2 to List All Users", "\nEnter 3 to Login", "\nEnter 4 to Exit")
		var ch int
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			var fname, lname, userid, pass string
			var bell, biba uint8
			fmt.Println("Enter first name:")
			fmt.Scanln(&fname)
			fmt.Println("Enter last Name:")
			fmt.Scanln(&lname)
			fmt.Println("Enter userid:")
			fmt.Scanln(&userid)
			fmt.Println("Enter password:")
			fmt.Scanln(&pass)
			fmt.Println("Enter your bell la padula level:")
			fmt.Scanln(&bell)
			fmt.Println("Enter your BIBA level:")
			fmt.Scanln(&biba)
			user.New(fname, lname, userid, pass, bell, biba)
			fmt.Println("User Add Successfully")
		case 2:
			user.ListAllUser()
		case 3:
			var userid, pass string
			fmt.Println("Enter userid:")
			fmt.Scanln(&userid)
			fmt.Println("Enter password:")
			fmt.Scanln(&pass)
			check, bell, biba := user.CheckUser(userid, pass)
			if check {
				fmt.Println("Login Successful")
				FileOperation(bell, biba)
			} else {
				fmt.Println("Login Unsuccessful")
			}
		case 4:
			user.WriteData()
			fmt.Println("Exiting")
			os.Exit(1)
		default:
			fmt.Println("Wrong Choice")
		}
	}
}

func FileOperation(bell, biba uint8) {
	files.Read()
	for {
		fmt.Println("\tWelcome", "\nEnter 1 to Read Files", "\nEnter 2 to Write Files", "\nEnter 3 to Exit")
		var ch int
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			fmt.Println("The Read able files are")
			cnt := files.ReadAble(bell, biba)
			if cnt <= 0 {
				fmt.Println("No files to read")
				continue
			}
			var name string
			fmt.Println("Enter any listed files name with extension:")
			fmt.Scanln(&name)
			err := files.ReadFile(name)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			fmt.Println("The Write able files are")
			cnt := files.WriteAble(bell, biba)
			if cnt <= 0 {
				fmt.Println("No files to write")
				continue
			}
			var name, phrase string
			fmt.Println("Enter any listed files name with extension:")
			fmt.Scanln(&name)
			fmt.Println("Enter any phrase:")
			fmt.Scanln(&phrase)
			err := files.WriteFile(name, phrase)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 3:
			return
		case 4:
			fmt.Println("Wrong Choice")
		}
	}
}
