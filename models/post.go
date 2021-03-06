package models

import (
	"fmt"
	"os"
	"strings"
)

type Post struct {
	Username string
	Text     string
}

func AddPostFile(p Post) (err error) {
	err = os.MkdirAll(fmt.Sprintf("./data/posts/%s/", p.Username), 0700)
	if err != nil {
		return
	}

	files, err := os.ReadDir(fmt.Sprintf("./data/posts/%s/", p.Username))
	if err != nil {
		return
	}

	err = os.WriteFile(fmt.Sprintf("./data/posts/%s/%d.txt", p.Username, len(files)+1), []byte(p.Text), 0600)
	if err != nil {
		return
	}

	return
}

func ShowPostsByUserName(name string) (err error) {
	files, err := os.ReadDir(fmt.Sprintf("./data/posts/%s/", name))
	if err != nil {
		return
	}

	for _, f := range files {
		fmt.Printf("Key: %s\n", strings.TrimSuffix(f.Name(), ".txt"))

		data, err := os.ReadFile(fmt.Sprintf("./data/posts/%s/%s", name, f.Name()))
		if err != nil {
			break
		}

		fmt.Printf("Content: %s\n", data)
		fmt.Println()
	}

	return
}
