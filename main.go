package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"taeho/acc/cli"
)

// Alias 옵션을 저장할 배열
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

	// 지정 가능한 서브커맨드
	// add: 계정 정보 추가
	// del: 계정 정보 삭제
	// ls: 계정 정보 조회
	// mod: 계정 정보 변경
	const addStr = "add"
	const delStr = "del"
	const lsStr = "ls"
	const modStr = "mod"

	// add 서브커맨드의 옵션들
	addCmd := flag.NewFlagSet(addStr, flag.ExitOnError)
	addT := addCmd.String("t", "", "Title")
	addU := addCmd.String("u", "", "Username")
	addP := addCmd.String("p", "", "Password")
	addL := addCmd.String("l", "", "Location")
	addE := addCmd.String("e", "", "Email")
	addM := addCmd.String("m", "", "Memo")
	addF := addCmd.String("f", "", "Insert from file")
	addCmd.Var(&as, "a", "Alias")

	// del 서브커맨드의 옵션들
	delCmd := flag.NewFlagSet(delStr, flag.ExitOnError)
	delT := delCmd.Bool("t", false, "Title")
	delU := delCmd.Bool("u", false, "Username")
	delP := delCmd.Bool("p", false, "Password")
	delL := delCmd.Bool("l", false, "Location")
	delE := delCmd.Bool("e", false, "Email")
	delM := delCmd.Bool("m", false, "Memo")
	delA := delCmd.Bool("a", false, "Alias")

	// ls 서브커맨드의 옵션들
	lsCmd := flag.NewFlagSet(lsStr, flag.ExitOnError)
	lsO := lsCmd.String("o", "short", "Output format [ short, wide, format= ]")
	lsI := lsCmd.Int("i", 0, "Search using index")
	lsT := lsCmd.String("t", "", "Search using title")
	lsU := lsCmd.String("u", "", "Search using username")
	lsCmd.Var(&as, "a", "Alias")

	// mod 서브커맨드의 옵션들
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
	// add 서브커맨드의 경우
	case addStr:
		addCmd.Parse(os.Args[2:])
		if len(*addF) > 0 {
			cli.AddFromFile(*addF)
		} else {
			cli.Add(*addT, *addU, *addP, *addL, *addE, *addM, as)
		}

	// del 서브커맨드의 경우
	case delStr:
		delCmd.Parse(os.Args[2:])
		indexStrSlice := delCmd.Args()
		indexSlice, err := convertIndex(indexStrSlice)
		if err != nil {
			log.Panicf(cli.PanicFormat, err)
		}
		cli.Del(indexSlice, *delT, *delU, *delP, *delL, *delE, *delM, *delA)

	// ls 서브커맨드의 경우
	case lsStr:
		lsCmd.Parse(os.Args[2:])
		cli.Ls(*lsI, *lsT, *lsU, as, *lsO)

	// mod 서브커맨드의 경우
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

// 문자열로 입력된 인덱스를 인트로 변환하는 함수
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
