package cmds

import (
	"fmt"
	"log"
	"syscall"

	"db"

	"golang.org/x/crypto/ssh/terminal"
)

func Login(username string, password string) {
	if username == "" {
		fmt.Print("Enter username: ")
		argn, err := fmt.Scanf("%s", &username)
		if err != nil {
			log.Panicln("Error: Failed to read username")
		} else if argn < 1 {
			log.Panicln("Error: You didn't enter username")
		}
	}

	if password == "" {
		fmt.Print("Enter password: ")
		bytePwd, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			log.Panicln("Error: Failed to read password")
		}
		password = string(bytePwd)
	}

	db.SetUsername(username)
	db.SetPassword(password)
	db.NewClient()
	db.PingConnection()
	db.DelClient()
}