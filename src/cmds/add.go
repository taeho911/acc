package cmds

import (
	"fmt"

	"db"
)

func Add(title, url, uid, pwd, email, alias, memo string) {
	doc := db.Acc{
		Index: 1,
		Title: title,
		Url: url,
		Uid: uid,
		Pwd: pwd,
		Email: email,
		Alias: []string{alias},
		Memo: memo,
	}

	fmt.Println(doc)
}