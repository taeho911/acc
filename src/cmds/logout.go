package cmds

import (
	"os"
	"log"
)

func Logout() {
	log.Println(os.Getenv("MONGO_USERNAME"))
	log.Println(os.Getenv("MONGO_PASSWORD"))
}