package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

// type Notification interface {
// 	Send(to, msg string)
// }

// type SmsNotification struct{}

// func (s *SmsNotification) Send(to, msg string) {
// 	val := fmt.Sprintf("SmsNotification %s sent to %s", msg, to)
// 	fmt.Println(val)
// }

// type EmailNotification struct{}

// func (e *EmailNotification) Send(to, msg string) {
// 	val := fmt.Sprintf("EmailNotification %s sent to %s", msg, to)
// 	fmt.Println(val)
// }

// func main() {
// 	fmt.Println("Hello, World!")
// 	to := "Raju"
// 	msg := "Hey how are you"

// 	var notif Notification
// 	notif = &EmailNotification{}
// 	notif.Send(to, msg)

// 	notif = &SmsNotification{}
// 	notif.Send(to, msg)
// }

/*
Let’s move to Phase 2 — some added requirements:

The system should support sending the same message over multiple channels at once, for example: both Email and SMS.

How would you extend your current design to support multi-channel notifications?

Feel free to walk me through the approach first before coding.
*/

// func main() {
// 	fmt.Println("Hello, World!")
// 	to := "Raju"
// 	msg := "Hey how are you"

// 	notif := &EmailNotification{}
// 	notif2 := &SmsNotification{}
// 	var sy sync.WaitGroup
// 	sy.Add(2)
// 	go func () {
// 		defer sy.Done();
// 		notif.Send(to, msg)
// 	}()
// 	go func () {
// 		defer sy.Done();
// 		notif2.Send(to, msg)
// 	}()
// 	sy.Wait()
// }

/*
Phase 3 — Abstractions & Extensibility:

Let’s say now you want to support:
  - Dynamically selecting which channels to send on (based on user preferences).
  - Sending to multiple users.
  - Logging failures (e.g., SMS failed, Email succeeded).
*/

type Notification interface {
	Send(to, msg string) error
}

type SmsNotification struct{}

func (s *SmsNotification) Send(to, msg string) error {
	if strings.TrimSpace(to) == "" {
		return errors.New("SMS: invalid recipient")
	}
	fmt.Printf("SMS sent to %s: %s\n", to, msg)
	return nil
}

type EmailNotification struct{}

func (e *EmailNotification) Send(to, msg string) error {
	if strings.TrimSpace(to) == "" {
		return errors.New("Email: invalid recipient")
	}
	fmt.Printf("Email sent to %s: %s\n", to, msg)
	return nil
}

type User struct {
	Name       string
	Channels   []Notification
}

type NotificationService struct{}

func (ns *NotificationService) SendToUsers(users []User, msg string) {
	var wg sync.WaitGroup

	for _, user := range users {
		user := user
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, notifier := range user.Channels {
				err := notifier.Send(user.Name, msg)
				if err != nil {
					fmt.Printf("Error sending to %s via %T: %v\n", user.Name, notifier, err)
				}
			}
		}()
	}

	wg.Wait()
}

func main() {
	email := &EmailNotification{}
	sms := &SmsNotification{}

	users := []User{
		{
			Name:     "Ram",
			Channels: []Notification{email, sms},
		},
		{
			Name:     "Saket",
			Channels: []Notification{sms},
		},
	}

	msg := "Hey, how are you?"
	service := &NotificationService{}
	service.SendToUsers(users, msg)
}