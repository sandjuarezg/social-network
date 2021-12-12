package friend

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func AddFriendFile(username, usernameFriend string) (err error) {
	if username == usernameFriend {
		err = errors.New("that's your username")
		return
	}

	_, err = os.Stat(fmt.Sprintf("./data/users/%s.txt", usernameFriend))
	if os.IsNotExist(err) {
		err = errors.New("username not found")
		return
	}

	var aux string

	if username > usernameFriend {
		aux = fmt.Sprintf("./data/friends/%s_%s.txt", username, usernameFriend)
	} else {
		aux = fmt.Sprintf("./data/friends/%s_%s.txt", usernameFriend, username)
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
	d := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())

	_, err = file.WriteString(fmt.Sprintf("friends since %s", d))
	if err != nil {
		err = errors.New("error to write")
		return
	}

	return
}

func GetFriendsByUserName(name string) (friends []string, err error) {
	files, err := os.ReadDir("./data/friends/")
	if err != nil {
		return
	}

	for _, file := range files {

		aux := strings.Split(strings.TrimSuffix(file.Name(), ".txt"), "_")
		var s string

		for i, v := range aux {

			if v != name {
				continue
			}

			if i == 0 {
				s = fmt.Sprintf("%s and you are ", aux[i+1])
			} else {
				s = fmt.Sprintf("%s and you are ", aux[i-1])
			}

			f, err := os.Open(fmt.Sprintf("./data/friends/%s", file.Name()))
			if err != nil {
				break
			}

			content, err := io.ReadAll(f)
			if err != nil {
				break
			}

			s += fmt.Sprintf("%s\n", content)
			friends = append(friends, s)

		}

	}

	return
}

func DeleteFriend(username, usernameFriend string) (err error) {
	var aux string

	if username > usernameFriend {
		aux = fmt.Sprintf("./data/friends/%s_%s.txt", username, usernameFriend)
	} else {
		aux = fmt.Sprintf("./data/friends/%s_%s.txt", usernameFriend, username)
	}

	err = os.Remove(aux)
	if err != nil {
		err = errors.New("username not found")
		return
	}

	return
}
