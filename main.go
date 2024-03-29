package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"taeho/acc/cli"
)

type aliasSlice []string

func (as *aliasSlice) String() string {
	return ""
}

func (as *aliasSlice) Set(val string) error {
	*as = append(*as, val)
	return nil
}

func main() {
	var as aliasSlice

	const addStr = "add"
	const delStr = "del"
	const lsStr = "ls"
	const modStr = "mod"

	addCmd := flag.NewFlagSet(addStr, flag.ExitOnError)
	addT := addCmd.String("t", "", "Title")
	addU := addCmd.String("u", "", "Username")
	addP := addCmd.String("p", "", "Password")
	addL := addCmd.String("l", "", "Location")
	addE := addCmd.String("e", "", "Email")
	addM := addCmd.String("m", "", "Memo")
	addCmd.Var(&as, "a", "Alias")

	delCmd := flag.NewFlagSet(delStr, flag.ExitOnError)
	delT := delCmd.Bool("t", false, "Title")
	delU := delCmd.Bool("u", false, "Username")
	delP := delCmd.Bool("p", false, "Password")
	delL := delCmd.Bool("l", false, "Location")
	delE := delCmd.Bool("e", false, "Email")
	delM := delCmd.Bool("m", false, "Memo")
	delA := delCmd.Bool("a", false, "Alias")

	lsCmd := flag.NewFlagSet(lsStr, flag.ExitOnError)
	lsO := lsCmd.String("o", "short", "Output format [ short, wide, format= ]")
	lsI := lsCmd.Int("i", 0, "Search using index")
	lsT := lsCmd.String("t", "", "Search using title")
	lsU := lsCmd.String("u", "", "Search using username")
	lsCmd.Var(&as, "a", "Alias")

	modCmd := flag.NewFlagSet(modStr, flag.ExitOnError)
	modT := modCmd.String("t", "", "Title")
	modU := modCmd.String("u", "", "Username")
	modP := modCmd.String("p", "", "Password")
	modL := modCmd.String("l", "", "location")
	modE := modCmd.String("e", "", "Email")
	modM := modCmd.String("m", "", "Memo")
	modCmd.Var(&as, "a", "Alias")
	modPush := modCmd.Bool("push", false, "Push alias")
	modPull := modCmd.Bool("pull", false, "Pull alias")

	if len(os.Args) < 2 {
		fmt.Println("Subcommands :: [add, del, ls, mod]")
		return
	}

	switch os.Args[1] {
	case addStr:
		addCmd.Parse(os.Args[2:])
		cli.Add(*addT, *addU, *addP, *addL, *addE, *addM, as)

	case delStr:
		delCmd.Parse(os.Args[2:])
		indexStrSlice := delCmd.Args()
		indexSlice, err := convertIndex(indexStrSlice)
		if err != nil {
			log.Panicf(cli.PanicFormat, err)
		}
		cli.Del(indexSlice, *delT, *delU, *delP, *delL, *delE, *delM, *delA)

	case lsStr:
		lsCmd.Parse(os.Args[2:])
		cli.Ls(*lsI, *lsT, *lsU, as, *lsO)

	case modStr:
		modCmd.Parse(os.Args[2:])
		indexStrSlice := modCmd.Args()
		indexSlice, err := convertIndex(indexStrSlice)
		if err != nil {
			log.Panicf(cli.PanicFormat, err)
		}
		if *modPush && *modPull {
			fmt.Println("Invalid flags :: Cannot set push/pull flag simultaneously")
			return
		}
		cli.Mod(indexSlice, *modT, *modU, *modP, *modL, *modE, *modM, as, *modPull, *modPush)

	default:
		fmt.Printf("Invalid subcommand :: %s\n", os.Args[1])
	}
}

func convertIndex(indexStrSlice []string) ([]int, error) {
	aliasSlice := make([]int, len(indexStrSlice))
	var err error
	for i, item := range indexStrSlice {
		if aliasSlice[i], err = strconv.Atoi(item); err != nil {
			return nil, err
		}
	}
	return aliasSlice, nil
}
