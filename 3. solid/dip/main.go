package main

import "fmt"

type Messenger interface {
    Send(user, message string)
}

type EmailSender struct {}

func (e *EmailSender) Send(user, message string) {
	fmt.Println("Sending email to", user, ":", message)
}

type SmsSender struct {}

func (e *SmsSender) Send(user, message string) {
	fmt.Println("Sending SMS to", user, ":", message)
}

// UserService depends directly on EmailSender (a concrete type), violating DIP.
type UserService struct {
	//emailSender *EmailSender
	Messenger
}

func (u *UserService) Notify(user, message string) {
	u.Messenger.Send(user, message)
}

func main() {
	fmt.Println("Hello, World!")
	es := &EmailSender{}
	us := &UserService{
		Messenger: es,
	}
	us.Notify("Ram", "Hey Ram")

	sms := &SmsSender{}
	us2 := &UserService{
		Messenger: sms,
	}
	us2.Notify("Ram", "Hey Ram")
}