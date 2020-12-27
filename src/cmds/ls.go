package cmds

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"

	"db"

	"go.mongodb.org/mongo-driver/bson"
)

func Ls(allFlag bool, id int, outputFormat, title, alias, uid string) {
	var result []bson.M

	if id != 0 || title != "" || alias != "" || uid != "" {
		result = db.Find(id, title, alias, uid)
	} else if allFlag == true {
		result = db.FindAll()
	} else {
		result = db.FindAll()
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
			printWide(result)
		}
	}
}

func printWide(result []bson.M) {
	for _, item := range result {
		fmt.Println("_id:\t", item["_id"])
		fmt.Println("title:\t", item["title"])
		fmt.Println("uid:\t", item["uid"])
		fmt.Println("pwd:\t", item["pwd"])
		fmt.Println("url:\t", item["url"])
		fmt.Println("email:\t", item["email"])
		fmt.Println("alias:\t", item["alias"])
		fmt.Println("memo:\t", item["memo"])
		fmt.Println()
	}
}

func printShort(result []bson.M) {
	fmt.Printf("TITLE\t\tUID\t\tPWD\n")
	for _, item := range result {
		fmt.Printf("%s\t\t%s\t\t%s\n", item["title"], item["uid"], item["pwd"])
	}
}

func convertFormat(format string, item bson.M) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, value := range item["alias"].(bson.A) {
		sb.WriteString(value.(string))
		if i == (len(item["alias"].(bson.A)) - 1) {
			sb.WriteString("]")
		} else {
			sb.WriteString(", ")
		}
	}

	// switch t := item["alias"].(type) {
	// case []string:
	// 	for i, item := range t {
	// 		fmt.Println("i:", i, "item:", item)
	// 	}
	// case bson.A:
	// 	fmt.Println("bson.A")
	// 	for i, item := range t {
	// 		fmt.Println("i:", i, "item:", item)
	// 	}
	// default:
	// 	fmt.Printf("type: %T\n", t)
	// }
	
	convertedString := format
	convertedString = strings.ReplaceAll(convertedString, "%i", strconv.FormatFloat(item["_id"].(float64), 'f', -1, 64))
	convertedString = strings.ReplaceAll(convertedString, "%t", item["title"].(string))
	convertedString = strings.ReplaceAll(convertedString, "%u", item["uid"].(string))
	convertedString = strings.ReplaceAll(convertedString, "%p", item["pwd"].(string))
	convertedString = strings.ReplaceAll(convertedString, "%U", item["url"].(string))
	convertedString = strings.ReplaceAll(convertedString, "%e", item["email"].(string))
	convertedString = strings.ReplaceAll(convertedString, "%a", sb.String())
	convertedString = strings.ReplaceAll(convertedString, "%m", item["memo"].(string))
	return convertedString
}