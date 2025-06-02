package main

import (
	"notificationSystem/notification"
	"notificationSystem/service"
)

func main() {
	email := &notification.EmailNotification{}
	sms := &notification.SmsNotification{}

	users := []notification.User{
		{Name: "Ram", Channels: []notification.Notification{email, sms}},
		{Name: "Saket", Channels: []notification.Notification{sms}},
	}

	msg := "Hey, how are you?"
	notifier := &service.NotificationService{}
	notifier.SendToUsers(users, msg)
}