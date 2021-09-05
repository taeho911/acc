package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"taeho/acc/api"

	"golang.org/x/term"
)

func Add(title, username, password, location, email, memo string, aliasSlice []string) {
	reader := bufio.NewReader(os.Stdin)
	if title == "" {
		fmt.Print("Title >> ")
		title, _ = reader.ReadString('\n')
	}
	if username == "" {
		fmt.Print("Username >> ")
		username, _ = reader.ReadString('\n')
	}
	if password == "" {
		fmt.Print("Password >> ")
		bytePwd, err1 := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err1 != nil {
			log.Panicf(PanicFormat, err1)
		}
		fmt.Print("Password Confirm >> ")
		bytePwdConfirm, err2 := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err2 != nil {
			log.Panicf(PanicFormat, err2)
		}
		password = string(bytePwd)
		passwordConfirm := string(bytePwdConfirm)
		if password != passwordConfirm {
			log.Panicf(PanicFormat, "Password confirmation failed")
		}
	}
	if location == "" {
		fmt.Print("Location >> ")
		location, _ = reader.ReadString('\n')
	}
	if email == "" {
		fmt.Print("Email >> ")
		email, _ = reader.ReadString('\n')
	}
	if memo == "" {
		fmt.Print("Memo >> ")
		memo, _ = reader.ReadString('\n')
	}
	if len(aliasSlice) < 1 {
		fmt.Print("Alias >> ")
		alias, _ := reader.ReadString('\n')
		aliasSlice = strings.Fields(alias)
	}

	title = strings.TrimSpace(title)
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	location = strings.TrimSpace(location)
	email = strings.TrimSpace(email)
	memo = strings.TrimSpace(memo)

	var acc api.Acc
	acc.Title = title
	acc.Username = username
	acc.Password = password
	acc.Location = location
	acc.Email = email
	acc.Memo = memo
	acc.Alias = aliasSlice

	result, err2 := api.InsertOne(acc)
	if err2 != nil {
		log.Panicf(PanicFormat, err2)
	}
	fmt.Println(result)
}

func AddFromFile(file string) {
	fileByte, err1 := os.ReadFile(file)
	if err1 != nil {
		log.Panicf(PanicFormat, err1)
	}
	var accs []api.Acc
	json.Unmarshal(fileByte, &accs)
	result, err2 := api.InsertMany(accs)
	if err2 != nil {
		log.Panicf(PanicFormat, err2)
	}
	for _, item := range result {
		fmt.Println(item)
	}
}
