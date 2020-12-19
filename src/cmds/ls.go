package cmds

import (
	"fmt"

	"db"

	"go.mongodb.org/mongo-driver/bson"
)

func Ls(a bool, o string) {
	var result []bson.M

	if a == true {
		result = db.FindAll()
	}

	switch o {
	case "wide":
		printWide(result)

	case "short":
		printShort(result)

	default:
		printWide(result)
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