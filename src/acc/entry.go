package main

import (
	"fmt"
	"os"
	"flag"

	"cmds"
)

func main() {
	loginCmd	:= flag.NewFlagSet("login", flag.ExitOnError)
	loginU		:= loginCmd.String("u", "", "Username")
	loginP		:= loginCmd.String("p", "", "Password")

	logoutCmd	:= flag.NewFlagSet("logout", flag.ExitOnError)

	addCmd		:= flag.NewFlagSet("add", flag.ExitOnError)
	addT		:= addCmd.String("t", "", "Title")
	addURL		:= addCmd.String("url", "", "URL")
	addU		:= addCmd.String("u", "", "User ID")
	addP		:= addCmd.String("p", "", "Password")
	addE		:= addCmd.String("e", "", "E-mail")
	addA		:= addCmd.String("a", "", "Alias")
	addM		:= addCmd.String("m", "", "Memo")
	
	delCmd		:= flag.NewFlagSet("del", flag.ExitOnError)

	lsCmd		:= flag.NewFlagSet("ls", flag.ExitOnError)
	lsA			:= lsCmd.Bool("a", true, "List up all")
	lsO			:= lsCmd.String("o", "wide", "Wide output")
	/*
	lsI			:= lsCmd.Int("i", 0, "Search using index")
	lsT			:= lsCmd.String("t", nil, "Search using index")

	modCmd		:= flag.NewFlagSet("mod", flag.ExitOnError)
	*/

	if len(os.Args) < 2 {
		fmt.Println("There is no subcommand")
		cmds.MyTest()
		return
	}

	switch os.Args[1] {
	case "login":
		loginCmd.Parse(os.Args[2:])
		cmds.Login(*loginU, *loginP)

	case "logout":
		logoutCmd.Parse(os.Args[2:])
		cmds.Logout()

	case "add":
		addCmd.Parse(os.Args[2:])
		cmds.Add(*addT, *addURL, *addU, *addP, *addE, *addA, *addM)

	case "del":
		delCmd.Parse(os.Args[2:])
		idArr := delCmd.Args()
		cmds.Del(idArr)

	case "ls":
		lsCmd.Parse(os.Args[2:])
		cmds.Ls(*lsA, *lsO)

	default:
		fmt.Printf("There is no %s subcommand", os.Args[1])
	}
}