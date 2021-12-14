package models

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func PreparePathDir() (err error) {
	err = os.MkdirAll("./data/users/", 0700)
	if err != nil {
		return
	}

	err = os.MkdirAll("./data/posts/", 0700)
	if err != nil {
		return
	}

	err = os.MkdirAll("./data/friends/", 0700)
	if err != nil {
		return
	}

	return
}

func CleanConsole() (err error) {
	time.Sleep(1 * time.Second)

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		return
	}

	return
}

func ScanTextWithMessagePrint(msg string) (text string, err error) {
	fmt.Println(msg)
	aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return
	}
	text = string(aux)

	return
}
