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
	addURL		:= addCmd.String("-url", "", "URL")
	addU		:= addCmd.String("u", "", "User ID")
	addP		:= addCmd.String("p", "", "Password")
	addE		:= addCmd.String("e", "", "E-mail")
	addA		:= addCmd.String("a", "", "Alias")
	addM		:= addCmd.String("m", "", "Memo")
	

	/*
	lsCmd		:= flag.NewFlagSet("ls", flag.ExitOnError)
	lsA			:= lsCmd.Bool("a", true, "List up all")
	lsAll		:= lsCmd.Bool("-all", true, "List up all")
	lsI			:= lsCmd.Int("i", 0, "Search using index")
	lsIndex		:= lsCmd.Int("-index", 0, "Search using index")
	lsT			:= lsCmd.String("t", nil, "Search using index")
	
	delCmd		:= flag.NewFlagSet("del", flag.ExitOnError)

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

	default:
		fmt.Printf("There is no %s subcommand", os.Args[1])
	}
}