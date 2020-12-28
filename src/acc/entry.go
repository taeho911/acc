package main

import (
	"fmt"
	"os"
	"flag"

	"cmds"
)

func main() {
	const loginStr = "login"
	const logoutStr = "logout"
	const addStr = "add"
	const delStr = "del"
	const lsStr = "ls"
	const modStr = "mod"

	loginCmd	:= flag.NewFlagSet(loginStr, flag.ExitOnError)
	loginU		:= loginCmd.String("u", "", "Username")
	loginP		:= loginCmd.String("p", "", "Password")

	logoutCmd	:= flag.NewFlagSet(logoutStr, flag.ExitOnError)

	addCmd		:= flag.NewFlagSet(addStr, flag.ExitOnError)
	addT		:= addCmd.String("t", "", "Title")
	addU		:= addCmd.String("u", "", "User ID")
	addP		:= addCmd.String("p", "", "Password")
	addURL		:= addCmd.String("U", "", "URL")
	addE		:= addCmd.String("e", "", "E-mail")
	addA		:= addCmd.String("a", "", "Alias")
	addM		:= addCmd.String("m", "", "Memo")
	
	delCmd		:= flag.NewFlagSet(delStr, flag.ExitOnError)

	lsCmd		:= flag.NewFlagSet(lsStr, flag.ExitOnError)
	lsAll		:= lsCmd.Bool("all", true, "List up all")
	lsODescription := `Output format
-o="format:xxxx" triggers customized format
	Resources:
	%i = Index
	%t = Title
	%u = User ID
	%p = Password
	%U = URL
	%e = E-mail
	%a = Alias
	%m = Memo`
	lsO			:= lsCmd.String("o", "short", lsODescription)
	lsI			:= lsCmd.Int("i", 0, "Search using _id")
	lsT			:= lsCmd.String("t", "", "Search using title")
	lsU			:= lsCmd.String("u", "", "Search using user ID")
	lsA			:= lsCmd.String("a", "", "Search using alias")

	modCmd		:= flag.NewFlagSet(modStr, flag.ExitOnError)
	modT		:= modCmd.String("t", "", "Title")
	modU		:= modCmd.String("u", "", "User ID")
	modP		:= modCmd.String("p", "", "Password")
	modURL		:= modCmd.String("U", "", "URL")
	modE		:= modCmd.String("e", "", "E-mail")
	modA		:= modCmd.String("a", "", "Alias")
	modDel		:= modCmd.Bool("del", false, "Delete alias")
	modAdd		:= modCmd.Bool("add", false, "Add alias")
	modM		:= modCmd.String("m", "", "Memo")

	if len(os.Args) < 2 {
		fmt.Println("There is no subcommand")
		return
	}

	switch os.Args[1] {
	case loginStr:
		loginCmd.Parse(os.Args[2:])
		cmds.Login(*loginU, *loginP)

	case logoutStr:
		logoutCmd.Parse(os.Args[2:])
		cmds.Logout()

	case addStr:
		addCmd.Parse(os.Args[2:])
		cmds.Add(*addT, *addURL, *addU, *addP, *addE, *addA, *addM)

	case delStr:
		delCmd.Parse(os.Args[2:])
		idArr := delCmd.Args()
		cmds.Del(idArr)

	case lsStr:
		lsCmd.Parse(os.Args[2:])
		cmds.Ls(*lsAll, *lsI, *lsO, *lsT, *lsA, *lsU)

	case modStr:
		modCmd.Parse(os.Args[2:])
		idArr := modCmd.Args()
		if *modDel == true && *modAdd == true {
			fmt.Println("You cannot set del flag and add flag simultaneously")
			break
		}
		cmds.Mod(idArr, *modT, *modU, *modP, *modURL, *modE, *modA, *modM, *modDel, *modAdd)

	default:
		fmt.Printf("There is no [%s] subcommand\n", os.Args[1])
	}
}