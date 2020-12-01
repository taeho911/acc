package cmds

import (
	"os"
	"log"

	"db"
)

func Logout() {
	log.Println(os.Getenv("MONGO_USERNAME"))
	log.Println(os.Getenv("MONGO_PASSWORD"))
	db.UnsetUsername()
	db.UnsetPassword()
}