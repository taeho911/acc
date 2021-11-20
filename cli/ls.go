package cli

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"taeho/acc/api"
)

func Ls(index int, title, username string, aliasSlice []string, outputFormat string) {
	var result []api.Acc
	var err error

	if index != 0 || title != "" || username != "" || len(aliasSlice) > 0 {
		result, err = api.Find(index, title, username, aliasSlice)
		if err != nil {
			log.Panicf(PanicFormat, err)
		}
	} else {
		result, err = api.FindAll()
		if err != nil {
			log.Panicf(PanicFormat, err)
		}
	}

	switch outputFormat {
	case "wide":
		printWide(result)

	case "short":
		printShort(result)

	default:
		if matchCheck, _ := regexp.MatchString(`^format=.*`, outputFormat); matchCheck {
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

func printWide(result []api.Acc) {
	format := "%-10v %-30v\n"
	for _, item := range result {
		fmt.Printf(format, "Index: ", item.Index)
		fmt.Printf(format, "Title: ", item.Title)
		fmt.Printf(format, "Username: ", item.Username)
		fmt.Printf(format, "Password: ", item.Password)
		fmt.Printf(format, "Location: ", item.Location)
		fmt.Printf(format, "Email: ", item.Email)
		fmt.Printf(format, "Memo: ", item.Memo)
		fmt.Printf(format, "Alias: ", sliceToString(item.Alias))
		fmt.Println("---")
	}
}

func printShort(result []api.Acc) {
	indexMax := 7
	titleMax := 7
	usernameMax := 6
	passwordMax := 6
	for _, item := range result {
		if length := len(item.Title) + 2; length > titleMax {
			titleMax = length
		}
		if length := len(item.Username) + 2; length > usernameMax {
			usernameMax = length
		}
		if length := len(item.Password) + 2; length > passwordMax {
			passwordMax = length
		}
	}
	format := fmt.Sprintf("%%-%dv%%-%dv%%-%dv%%-%dv\n", indexMax, titleMax, usernameMax, passwordMax)
	if len(result) > 0 {
		fmt.Printf(format, "INDEX", "TITLE", "USER", "PASS")
		for _, item := range result {
			fmt.Printf(format, item.Index, item.Title, item.Username, item.Password)
		}
	}
}

func convertFormat(format string, item api.Acc) string {
	aliasStr := sliceToString(item.Alias)
	convertedString := format
	convertedString = strings.ReplaceAll(convertedString, "%i", strconv.Itoa(item.Index))
	convertedString = strings.ReplaceAll(convertedString, "%t", item.Title)
	convertedString = strings.ReplaceAll(convertedString, "%u", item.Username)
	convertedString = strings.ReplaceAll(convertedString, "%p", item.Password)
	convertedString = strings.ReplaceAll(convertedString, "%l", item.Location)
	convertedString = strings.ReplaceAll(convertedString, "%e", item.Email)
	convertedString = strings.ReplaceAll(convertedString, "%m", item.Memo)
	convertedString = strings.ReplaceAll(convertedString, "%a", aliasStr)
	return convertedString
}

func sliceToString(slice []string) string {
	var sb strings.Builder
	length := len(slice)
	if len(slice) > 0 {
		sb.WriteString("[ ")
		for i, val := range slice {
			sb.WriteString(val)
			sb.WriteString(" ")
			if i == length-1 {
				sb.WriteString("]")
			}
		}
	}
	return sb.String()
}
