package main

import (
	"fmt"
	"sync"
)

type logger struct{}

func (l *logger) log(message string) {
	fmt.Println("logging... ", message)
}

var (
	logr     *logger
	//once     sync.Once
	logrLock sync.Mutex
)


// The sync.Once ensures that the initialization code inside the Do method is executed only once,
// even in the presence of multiple goroutines. It provides thread-safe lazy initialization.
// func GetLoggerInstance() *logger {
// 	once.Do(func() {
// 		logr = &logger{}
// 	})
// 	return logr
// }

func GetLoggerInstance() *logger {
	logrLock.Lock()
	defer logrLock.Unlock()

	if logr == nil {
		logr = &logger{}
	}
	return logr
}

func main() {
	logger1 := GetLoggerInstance()
	logger2 := GetLoggerInstance()
	fmt.Println(logger1 == logger2) // true

	logger1.log("This is the first log")
	logger2.log("This is the second log")
}
