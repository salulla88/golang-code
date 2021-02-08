//Challenge call github api and fetch user's name and count of public repos
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

func UserInfo(login string) (*User, error) {
	resp, error := http.Get("https://api.github.com/users/" + login)
	if error != nil {
		log.Fatalf("failed to fetch information from github : %s\n", error)
		return nil, error
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	user := &User{}
	if err := dec.Decode(user); err != nil {
		log.Fatalf("error: cant decode - %s", err)
	}

	return user, nil
}

func main() {
	log.Println("==========Start=========")

	user, err := UserInfo("sandeeplulla")

	if err != nil {
		log.Fatalf("error : %s", err)
	}

	fmt.Printf("%+v\n", user)

	log.Println("==========End=========")
}
