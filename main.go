package main

import (
	"fmt"
	"log"

	"github.com/sandjuarezg/social-network/packages/friend"
	"github.com/sandjuarezg/social-network/packages/functionality"
	"github.com/sandjuarezg/social-network/packages/post"
	"github.com/sandjuarezg/social-network/packages/user"
)

func main() {
	var (
		opc  int
		exit bool
	)

	err := functionality.PreparePathDir()
	if err != nil {
		log.Fatalln(err)
	}

	for !exit {

		err := functionality.CleanConsole()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("0. Exit")
		fmt.Println("----------")
		fmt.Println("1. Log in")
		fmt.Println("2. Sing up")
		fmt.Scanln(&opc)

		err = functionality.CleanConsole()
		if err != nil {
			log.Println(err)
			continue
		}

		switch opc {
		case 0:

			exit = true
			fmt.Println(". . . .  B Y E  . . . .")

			err := functionality.CleanConsole()
			if err != nil {
				log.Println(err)
				continue
			}

		case 1:

			var back bool

			name, err := functionality.ScanTextByMessage("Enter user name")
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println()
			passwd, err := functionality.ScanTextByMessage("Enter user password")
			if err != nil {
				log.Println(err)
				continue
			}

			u, err := user.LogIn(name, passwd)
			if err != nil {
				log.Println(err)
				continue
			}

			for !back {

				err := functionality.CleanConsole()
				if err != nil {
					log.Println(err)
					continue
				}

				opc = 0
				fmt.Printf("~ Welcome %s ~\n", u.Name)
				fmt.Println("0. Sign off")
				fmt.Println("1. Delete account")
				fmt.Println("-------------------")
				fmt.Println("2. Add post")
				fmt.Println("3. Add friend")
				fmt.Println("4. Delete friend")
				fmt.Println("-------------------")
				fmt.Println("5. Show your posts")
				fmt.Println("6. Show your friends")
				fmt.Scanln(&opc)

				err = functionality.CleanConsole()
				if err != nil {
					log.Println(err)
					continue
				}

				switch opc {
				case 0:

					back = true

					err := functionality.CleanConsole()
					if err != nil {
						log.Println(err)
						continue
					}

				case 1:
					var opc string

					fmt.Println("Are you sure you want to delete this account? y/n")
					fmt.Scanln(&opc)

					if opc != "y" {
						continue
					}

					err = u.DeleteAccount()
					if err != nil {
						log.Println(err)
						continue
					}

					back = true

					fmt.Println()
					fmt.Println("Account deleted successfully")

				case 2:

					text, err := functionality.ScanTextByMessage("Enter post text")
					if err != nil {
						log.Println(err)
						continue
					}

					err = post.AddPostFile(post.Post{Username: u.Name, Text: text})
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Println("Post added successfully")

				case 3:

					name, err = functionality.ScanTextByMessage("Enter username")
					if err != nil {
						log.Println(err)
						continue
					}

					err = friend.AddFriendFile(u.Name, name)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Printf("%s now is your friend\n", name)

				case 4:

					name, err = functionality.ScanTextByMessage("Enter username")
					if err != nil {
						log.Println(err)
						continue
					}

					err = friend.DeleteFriend(u.Name, name)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Printf("%s now isn't your friend\n", name)

				case 5:

					err = post.ShowPostsByUserName(u.Name)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println("Press ENTER to continue")
					fmt.Scanln()

				case 6:

					friends, err := friend.GetFriendsByUserName(u.Name)
					if err != nil {
						log.Println(err)
						continue
					}

					if len(friends) == 0 {
						fmt.Println("you don't have friends yet")
						continue
					}

					for _, v := range friends {
						fmt.Println(v)
					}

					fmt.Println("Press ENTER to continue")
					fmt.Scanln()

				default:

					fmt.Println("Option not valid")

				}
			}

		case 2:

			name, err := functionality.ScanTextByMessage("Enter user name")
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println()
			passwd, err := functionality.ScanTextByMessage("Enter user password")
			if err != nil {
				log.Println(err)
				continue
			}

			err = user.AddUserFile(user.User{Name: name, Passwd: passwd})
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println()
			fmt.Println("User added successfully")

		default:

			fmt.Println("Option not valid")

		}

	}

}
