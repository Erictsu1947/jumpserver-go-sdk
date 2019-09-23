package main

import (
	"fmt"
	"strings"

	jms "github.com/Erictsu1947/jumpserver-go-sdk"
)

func main() {
	//r := bufio.NewReader(os.Stdin)

	//fmt.Print("JumpServer URL: ")
	//jiraURL, _ := r.ReadString('\n')
	//
	//fmt.Print("JumpServer Username: ")
	//username, _ := r.ReadString('\n')
	//
	//fmt.Print("JumpServer Password: ")
	//bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	//password := string(bytePassword)

	jmsURL := "http://localhost"
	authURL := "http://localhost/api/users/v1/auth/"
	username := "admin"
	password := "admin"

	tp := jms.TokenAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
		AuthURL: strings.TrimSpace(authURL),
	}

	client, err := jms.NewClient(tp.Client(), strings.TrimSpace(jmsURL))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	// search
	matchUserList, _, err := client.Users.Search("admin")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	for _, user := range matchUserList {
		fmt.Println(user.Email)
	}

	// list
	userList, _, err := client.Users.GetList()
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
	for _, user := range userList {
		fmt.Println(user.Email)
	}


}
