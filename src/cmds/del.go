package cmds

import (
	"fmt"
	"strconv"

	"db"
)

func Del(idArr []string) {
	idIntArr := make([]int, len(idArr))
	var err error = nil
	for i, item := range idArr {
		if idIntArr[i], err = strconv.Atoi(item); err != nil {
			fmt.Println(item, "cannot be converted to int")
		}
	}

	result := db.DeleteMany(idIntArr)
	fmt.Println("# Deleted documents:", result.DeletedCount)
}