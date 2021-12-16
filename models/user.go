package models

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name   string
	Passwd string
}

func ExistUserByUsername(name string) (ban bool, err error) {
	_, err = os.Stat(fmt.Sprintf("./data/users/%s.txt", name))
	if err != nil {
		if os.IsNotExist(err) {
			return
		}

		return
	}

	ban = true

	return
}

func AddUserFile(u User) (err error) {
	if strings.Contains(u.Name, "_") {
		err = errors.New("can't include the character '_'")
		return
	}

	ban, err := ExistUserByUsername(u.Name)
	if err != nil {
		return
	}

	if ban {
		err = errors.New("this username already exists")
		return
	}

	err = os.WriteFile(fmt.Sprintf("./data/users/%s.txt", u.Name), []byte(u.Passwd), 0600)
	if err != nil {
		err = errors.New("error to add user")
		return
	}

	return
}

func LogIn(name, passwd string) (u User, err error) {
	ban, err := ExistUserByUsername(name)
	if err != nil {
		return
	}

	if !ban {
		err = errors.New("user not found")
		return
	}

	content, err := os.ReadFile(fmt.Sprintf("./data/users/%s.txt", name))
	if err != nil {
		err = errors.New("error to read password")
		return
	}

	if string(content) != passwd {
		err = errors.New("incorrect password")
		return
	}

	u.Name = name

	return
}

func (u User) DeleteAccount() (err error) {
	friends, err := GetFriendsByUsername(u.Name)
	if err != nil {
		return
	}

	if len(friends) != 0 {
		err = errors.New("account can't be deleted due to data dependency")
		return
	}

	err = os.Remove(fmt.Sprintf("./data/users/%s.txt", u.Name))
	if err != nil {
		return
	}

	err = os.RemoveAll(fmt.Sprintf("./data/posts/%s", u.Name))
	if err != nil {
		return
	}

	return
}
