package main

import (
	"fmt"
)

type User struct {
	name string
	surname string
	age int
	currency int
}

func (u User) userInfoMethod(currency int) (string, int) {
	return u.name + " " + u.surname, currency + 100
}

func userInfoFunction(u User) User {
	u.name += "AAA"
	u.surname += "BBB"
	u.age += 5
	u.currency += 100

	return u
}

func main() {
	user := User {
		name: "Igor",
		surname: "Ganzage",
		age: 20,
		currency: 4000,
	}

	userA, userA_ := user.userInfoMethod(5000)
	userB := userInfoFunction(user)

	fmt.Println(userA, userA_)
	fmt.Println(userB)
}