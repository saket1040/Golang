package main

import "fmt"

type MessageSender interface{
	SendMessage(to string, message string)
}

type BasicSender struct{}

func (b *BasicSender) SendMessage(to string, message string) {
	fmt.Println("Sending plain message to", to, message)
}

type EncryptedSender struct {
	Wrapped MessageSender
}

func (e *EncryptedSender) SendMessage(to string, message string) {
	// Add encryption logic (simulated here)
	encrypted := "[ENCRYPTED] " + message
	fmt.Println("Encrypting message...")
	e.Wrapped.SendMessage(to, encrypted)
}

type LoggingSender struct {
	Wrapped MessageSender
}

func (l *LoggingSender) SendMessage(to string, message string) {
	fmt.Println("Logging: Sending message to", to)
	l.Wrapped.SendMessage(to, message)
}

func main() {
	sender := &LoggingSender{
    Wrapped: &EncryptedSender{
        Wrapped: &BasicSender{},
		},
	}
	sender.SendMessage("user", "hello")
}