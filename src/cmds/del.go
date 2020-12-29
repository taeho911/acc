package cmds

import (
	"fmt"
	"strconv"

	"db"
)

func Del(idArr []string, title, uid, pwd, url, email, alias, memo bool) {
	idIntArr := make([]int, len(idArr))
	var err error = nil
	for i, item := range idArr {
		if idIntArr[i], err = strconv.Atoi(item); err != nil {
			fmt.Println(item, "cannot be converted to int")
		}
	}
	
	if title == true || uid == true || pwd == true || url == true || email == true || alias == true || memo == true {
		result := db.EmptyFields(idIntArr, title, uid, pwd, url, email, alias, memo)
		fmt.Println("# Emptied documents:", result.ModifiedCount)
	} else {
		result := db.DeleteMany(idIntArr)
		fmt.Println("# Deleted documents:", result.DeletedCount)
	}
}