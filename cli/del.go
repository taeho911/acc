package cli

import (
	"fmt"
	"log"

	"taeho/acc/api"
)

func Del(indexSlice []int, title, username, password, location, email, memo, alias bool) {
	if title || username || password || location || email || memo || alias {
		result, err1 := api.EmptyFields(indexSlice, title, username, password, location, email, memo, alias)
		if err1 != nil {
			log.Panicf(PanicFormat, err1)
		}
		fmt.Println("Modified:", result)
	} else {
		result, err2 := api.DeleteMany(indexSlice)
		if err2 != nil {
			log.Panicf(PanicFormat, err2)
		}
		fmt.Println("Deleted:", result)
	}
}
