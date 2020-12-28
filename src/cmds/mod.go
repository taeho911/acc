package cmds

import (
	"fmt"
	"strconv"

	"db"
)

func Mod(idArr []string, title, uid, pwd, url, email, alias, memo string, delFlag, addFlag bool) {
	idIntArr := make([]int, len(idArr))
	var err error = nil
	for i, item := range idArr {
		if idIntArr[i], err = strconv.Atoi(item); err != nil {
			fmt.Println(item, "cannot be converted to int")
		}
	}

	result := db.UpdateMany(idIntArr, title, uid, pwd, url, email, alias, memo, delFlag, addFlag)
	fmt.Println("# Modified documents:", result.ModifiedCount)
}