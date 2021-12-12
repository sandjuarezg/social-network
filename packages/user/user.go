package user

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sandjuarezg/social-network/packages/friend"
)

type User struct {
	Name   string
	Passwd string
}

func AddUserFile(u User) (err error) {
	_, err = os.Stat(fmt.Sprintf("./data/users/%s.txt", u.Name))
	if !os.IsNotExist(err) {
		err = errors.New("this username already exists")
		return
	}

	if strings.Contains(u.Name, "_") {
		err = errors.New("can't include the character '_'")
		return
	}

	file, err := os.Create(fmt.Sprintf("./data/users/%s.txt", u.Name))
	if err != nil {
		err = errors.New("error to add user")
		return
	}
	defer file.Close()

	_, err = file.WriteString(u.Passwd)
	if err != nil {
		err = errors.New("error to write password")
		return
	}

	return
}

func LogIn(name, passwd string) (u User, err error) {
	file, err := os.Open(fmt.Sprintf("./data/users/%s.txt", name))
	if err != nil {
		err = errors.New("user not found")
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
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
	friends, err := friend.GetFriendsByUserName(u.Name)
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
