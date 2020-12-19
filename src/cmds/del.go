package cmds

import (
	"log"
	"strconv"

	"db"
)

func Del(idArr []string) {
	log.Println(idArr)

	idIntArr := make([]int, len(idArr))
	var err error = nil

	for i, item := range idArr {
		log.Println("item:", item)
		if idIntArr[i], err = strconv.Atoi(item); err != nil {
			log.Println(item, "cannot be converted to int")
		}
	}

	result := db.DeleteMany(idIntArr)
	if result == nil {
		log.Println("Failed to delete credentials")
	} else {
		log.Printf("Succeed to delete %v\n", result.DeletedCount)
	}
}