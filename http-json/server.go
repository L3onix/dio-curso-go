package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Users struct {
	users []User `json: "users"`
}

type User struct {
	name   string `json: "name"`
	_type  string `json: "type"`
	age    int    `json: "age"`
	social Social `json: "social"`
}

type Social struct {
	facebook string `json: "facebook"`
	twitter  string `json: "twitter"`
}

func main() {
	jsonFile, err := os.Open("./http-json/user.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var users Users
	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.users); i++ {
		fmt.Println("user type: " + users.users[i]._type)
		fmt.Println("user age: " + strconv.Itoa(users.users[i].age))
		fmt.Println("user name: " + users.users[i].name)
	}
}
