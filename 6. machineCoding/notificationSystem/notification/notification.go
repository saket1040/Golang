package notification

import (
	"errors"
	"fmt"
	"strings"
)

// Notification is the interface all channels implement.
type Notification interface {
	Send(to, msg string) error
}

// SmsNotification implements Notification
type SmsNotification struct{}

func (s *SmsNotification) Send(to, msg string) error {
	if strings.TrimSpace(to) == "" {
		return errors.New("SMS: invalid recipient")
	}
	fmt.Printf("SMS sent to %s: %s\n", to, msg)
	return nil
}

// EmailNotification implements Notification
type EmailNotification struct{}

func (e *EmailNotification) Send(to, msg string) error {
	if strings.TrimSpace(to) == "" {
		return errors.New("Email: invalid recipient")
	}
	fmt.Printf("Email sent to %s: %s\n", to, msg)
	return nil
}