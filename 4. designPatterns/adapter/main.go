package main

import "fmt"

// Third-party logging lib (you can't change this)
type ThirdPartyLogger struct{}

func (t *ThirdPartyLogger) WriteLog(msg string, severity int) {
	fmt.Printf("[Severity %d]: %s\n", severity, msg)
}

// Our expected interface
type Logger interface {
	Log(message string)
}

type LoggerAdapter struct {
	ThirdParty *ThirdPartyLogger
}

func (a *LoggerAdapter) Log(message string) {
	// Default severity = 1
	a.ThirdParty.WriteLog(message, 1)
}

func main() {
	// Use the adapter to wrap the third-party logger
	logger := &LoggerAdapter{
		ThirdParty: &ThirdPartyLogger{},
	}

	logger.Log("Something went wrong")
}