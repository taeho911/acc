package cmds

import (
	"fmt"
	"syscall"

	"db"

	"golang.org/x/crypto/ssh/terminal"
)

func Login(username string, password string) {
	if username == "" {
		fmt.Print("Enter username: ")
		argn, err := fmt.Scanf("%s", &username)
		if err != nil {
			fmt.Println("Error: Failed to read username")
			return
		} else if argn < 1 {
			fmt.Println("Error: You didn't enter username")
			return
		}
	}

	if password == "" {
		fmt.Print("Enter password: ")
		bytePwd, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			fmt.Println("Error: Failed to read password")
			return
		}
		password = string(bytePwd)
	}

	db.Ping()
}