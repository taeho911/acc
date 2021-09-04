package cli

import (
	"fmt"
	"log"

	"taeho/acc/api"
)

func Mod(indexSlice []int, title, username, password, location, email, memo string, aliasSlice []string, aliasPull, aliasPush bool) {
	var acc api.Acc
	acc.Title = title
	acc.Username = username
	acc.Password = password
	acc.Location = location
	acc.Email = email
	acc.Memo = memo
	acc.Alias = aliasSlice

	result, err1 := api.UpdateMany(indexSlice, acc, aliasPull, aliasPush)
	if err1 != nil {
		log.Panicf(PanicFormat, err1)
	}
	fmt.Println("Modified:", result)
}
