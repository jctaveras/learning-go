package main

import (
	"fmt"
	"strings"
)

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	user := &person{
		firstName: "Jean Carlos",
		lastName: "Taveras",
		contactInfo: contactInfo{
			email: "jctaveras@intellisys.com.do",
			zipCode: 40003,
		},
	}
	
	fmt.Println(user.fullName())
	user.updateFirstName("Jyan")
	fmt.Println(user.fullName())

}

func (p person) fullName() string {
	return strings.Join([]string{p.firstName, p.lastName}, " ")
}

func (p *person) updateFirstName(name string) {
	(*p).firstName = name
}
