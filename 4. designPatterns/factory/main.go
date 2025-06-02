package main

import "fmt"

type Notifier interface {
	Notify(user string, message string)
}

type EmailNotifier struct{}

func (e *EmailNotifier) Notify(user string, message string) {
	fmt.Println("Sending email to", user, ":", message)
}

type SMSNotifier struct{}

func (s *SMSNotifier) Notify(user string, message string) {
	fmt.Println("Sending SMS to", user, ":", message)
}

type PushNotifier struct{}

func (p *PushNotifier) Notify(user string, message string) {
	fmt.Println("Sending Notification to", user, ":", message)
}

func GetNotifier(notificationType string) Notifier {
	switch notificationType {
	case "email":
		return &EmailNotifier{}
	case "sms":
		return &SMSNotifier{}
	case "push":
		return &PushNotifier{}
	}
	
	return nil
}

func main() {
	notifier := GetNotifier("email")
	notifier.Notify("user@example.com", "Welcome!")

	notifier = GetNotifier("sms")
	notifier.Notify("+1234567890", "Your OTP is 123456")

	notifier = GetNotifier("push")
	notifier.Notify("user123", "You have a new follower!")
}
