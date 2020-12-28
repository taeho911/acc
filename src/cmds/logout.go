package cmds

import (
	"os"
	"fmt"
)

func Logout() {
	fmt.Println(os.Getenv("MONGO_USERNAME"))
	fmt.Println(os.Getenv("MONGO_PASSWORD"))
}