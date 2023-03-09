package main

import (
	"fmt"
)

type Telegram struct {
	Username string
	Phone    string
}

func (t Telegram) Call() {
	fmt.Println("Telegram to", t.Username)
}

type Whatsapp struct {
	Phone string
}

func (w Whatsapp) Call() {
	fmt.Println("Whatsapp to", w.Phone)
}

type Instagram struct {
	Username string
	Email    string
}

type Caller interface {
	Call()
}

type Messenger interface {
	Message()
}

func main() {
	telegram := Telegram{
		Username: "paraparadox",
		Phone:    "987654321",
	}
	whatsapp := Whatsapp{
		Phone: "999888777",
	}
	//instagram := Instagram{
	//	Username: "paraparadox",
	//	Email:    "paraparadox@protonmail.com",
	//}

	callToFriend(telegram)
	callToFriend(whatsapp)
	//callToFriend(instagram)
}

func callToFriend(caller Caller) {
	fmt.Println("Calling to a friend through")
	caller.Call()
}

func messageToFriend(messenger Messenger) {
	fmt.Println("Messaging to a friend through")
	messenger.Message()
}
