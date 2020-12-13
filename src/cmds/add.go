package cmds

import (
	"os"
	"fmt"
	"syscall"
	"log"
	"bufio"
	"strings"

	"db"

	"golang.org/x/crypto/ssh/terminal"
)

func Add(title, url, uid, pwd, email, alias, memo string) {
	reader := bufio.NewReader(os.Stdin)
	
	if title == "" {
		fmt.Print("Title: ")
		title, _ = reader.ReadString('\n')
	}
	if uid == "" {
		fmt.Print("ID: ")
		uid, _ = reader.ReadString('\n')
	}
	if pwd == "" {
		fmt.Print("Password: ")
		bytePwd, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			log.Panicln("Error: Failed to read password")
		}
		pwd = string(bytePwd)
	}
	if url == "" {
		fmt.Print("URL: ")
		url, _ = reader.ReadString('\n')
	}
	if email == "" {
		fmt.Print("E-mail: ")
		email, _ = reader.ReadString('\n')
	}
	if alias == "" {
		fmt.Print("Alias: ")
		alias, _ = reader.ReadString('\n')
	}
	if memo == "" {
		fmt.Print("Memo: ")
		memo, _ = reader.ReadString('\n')
	}

	title = strings.TrimSpace(title)
	uid = strings.TrimSpace(uid)
	pwd = strings.TrimSpace(pwd)
	url = strings.TrimSpace(url)
	email = strings.TrimSpace(email)
	aliasArr := strings.Fields(alias)
	memo = strings.TrimSpace(memo)

	result := db.InsertOne(title, url, uid, pwd, email, memo, aliasArr)
	if result == nil {
		log.Println("Failed to post credential")
	}
}