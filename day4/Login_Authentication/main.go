package main

import (
	"fmt"
	"login/users"
	"os"
)

func main() {
	users.ReadData()
	for {
		fmt.Println("\tMenu", "\nEnter 1 to Add User", "\nEnter 2 to List All Users", "\nEnter 3 to Login", "\nEnter 4 to Exit")
		var ch int
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			var fname, lname, userid, pass string
			fmt.Println("Enter first name:")
			fmt.Scanln(&fname)
			fmt.Println("Enter last Name:")
			fmt.Scanln(&lname)
			fmt.Println("Enter userid:")
			fmt.Scanln(&userid)
			fmt.Println("Enter password:")
			fmt.Scanln(&pass)
			users.New(fname, lname, userid, pass)
			fmt.Println("User Add Successfully")
		case 2:
			users.ListAllUser()
		case 3:
			var userid, pass string
			fmt.Println("Enter userid:")
			fmt.Scanln(&userid)
			fmt.Println("Enter password:")
			fmt.Scanln(&pass)
			check := users.CheckUser(userid, pass)
			if check {
				fmt.Println("Login Successful")
			} else {
				fmt.Println("Login Unsuccessful")
			}
		case 4:
			users.WriteData()
			fmt.Println("Exiting")
			os.Exit(1)
		default:
			fmt.Println("Wrong Choice")
		}
	}
}
