package models

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Friend struct {
	FirtName   string
	SecondName string
	Date       string
}

func AddFriendFile(frds Friend) (err error) {
	if frds.FirtName == frds.SecondName {
		err = errors.New("that's your username")
		return
	}

	if !ExistUserByUsername(frds.FirtName) || !ExistUserByUsername(frds.SecondName) {
		err = errors.New("username not found")
		return
	}

	var aux string = fmt.Sprintf("./data/friends/%s_%s.txt", frds.SecondName, frds.FirtName)

	if frds.FirtName > frds.SecondName {
		aux = fmt.Sprintf("./data/friends/%s_%s.txt", frds.FirtName, frds.SecondName)
	}

	_, err = os.Stat(aux)
	if !os.IsNotExist(err) {
		err = errors.New("you are already friends")
		return
	}

	file, err := os.Create(aux)
	if err != nil {
		return
	}
	defer file.Close()

	t := time.Now()

	_, err = file.WriteString(fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day()))
	if err != nil {
		err = errors.New("error to write")
		return
	}

	return
}

func GetFriendsByUsername(name string) (frds []Friend, err error) {
	files, err := os.ReadDir("./data/friends/")
	if err != nil {
		return
	}

	for _, file := range files {

		aux := strings.Split(strings.TrimSuffix(file.Name(), ".txt"), "_")

		for _, v := range aux {

			if v != name {
				continue
			}

			var (
				auxFriend Friend
				content   []byte
			)

			content, err = os.ReadFile(fmt.Sprintf("./data/friends/%s", file.Name()))
			if err != nil {
				return
			}

			auxFriend.Date = string(content)

			for _, item := range aux {
				if item != name {
					auxFriend.SecondName = item
					break
				}
			}

			frds = append(frds, auxFriend)

		}

	}

	return
}

func DeleteFriend(username, usernameFriend string) (err error) {
	var aux string = fmt.Sprintf("./data/friends/%s_%s.txt", usernameFriend, username)

	if username > usernameFriend {
		aux = fmt.Sprintf("./data/friends/%s_%s.txt", username, usernameFriend)
	}

	err = os.Remove(aux)
	if err != nil {
		err = errors.New("username not found")
		return
	}

	return
}
