package service

import (
	"fmt"
	"notificationSystem/notification"
	"sync"
)

type NotificationService struct{}

func (ns *NotificationService) SendToUsers(users []notification.User, msg string) {
	var wg sync.WaitGroup

	for _, user := range users {
		user := user // capture range variable
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