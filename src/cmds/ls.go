package cmds

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
	"os"
	"text/tabwriter"

	"db"
)

func Ls(allFlag bool, id int, outputFormat, title, alias, uid string) {
	var result []db.Account

	if allFlag == true && id == 0 && title == "" && alias == "" && uid == "" {
		result = db.FindAll()
	} else {
		result = db.Find(id, title, alias, uid)
	}

	switch outputFormat {
	case "wide":
		printWide(result)

	case "short":
		printShort(result)

	default:
		if matchCheck, _ := regexp.MatchString(`^format:.*`, outputFormat); matchCheck {
			format := outputFormat[7:]
			for _, item := range result {
				convertedString := convertFormat(format, item)
				fmt.Println(convertedString)
			}
		} else {
			printShort(result)
		}
	}
}

func printWide(result []db.Account) {
	for _, item := range result {
		fmt.Println("_id:\t", item.Id)
		fmt.Println("title:\t", item.Title)
		fmt.Println("uid:\t", item.Uid)
		fmt.Println("pwd:\t", item.Pwd)
		fmt.Println("url:\t", item.Url)
		fmt.Println("email:\t", item.Email)
		fmt.Println("alias:\t", item.Alias)
		fmt.Println("memo:\t", item.Memo)
		fmt.Println()
	}
}

func printShort(result []db.Account) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', 0)
	fmt.Fprintf(w, "INDEX\tTITLE\tUID\tPWD\n")
	for _, item := range result {
		fmt.Fprintf(w, "%g\t%s\t%s\t%s\n", item.Id, item.Title, item.Uid, item.Pwd)
	}
	w.Flush()
}

func convertFormat(format string, item db.Account) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, value := range item.Alias {
		sb.WriteString(value)
		if i == (len(item.Alias) - 1) {
			sb.WriteString("]")
		} else {
			sb.WriteString(" ")
		}
	}

	convertedString := format
	convertedString = strings.ReplaceAll(convertedString, "%i", strconv.FormatFloat(item.Id, 'f', -1, 64))
	convertedString = strings.ReplaceAll(convertedString, "%t", item.Title)
	convertedString = strings.ReplaceAll(convertedString, "%u", item.Uid)
	convertedString = strings.ReplaceAll(convertedString, "%p", item.Pwd)
	convertedString = strings.ReplaceAll(convertedString, "%U", item.Url)
	convertedString = strings.ReplaceAll(convertedString, "%e", item.Email)
	convertedString = strings.ReplaceAll(convertedString, "%a", sb.String())
	convertedString = strings.ReplaceAll(convertedString, "%m", item.Memo)
	return convertedString
}