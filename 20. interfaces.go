package main

import "fmt"

// Interface - grouping of methods with given name, in this example common function for both structs
type ilogger interface {
	log(msg string, isError bool)
}

func logInfo(logger ilogger, msg string) {
	logger.log(msg, false)
}
func logError(logger ilogger, msg string) {
	logger.log(msg, true)
}

// Structs containing function with same signature, matching ilogger interface
type fileLogger struct {
	filePath string
}

func (logger fileLogger) log(msg string, isError bool) {
	// fake implementation for demo
	var level string = "Info"
	if isError {
		level = "Error"
	}
	fmt.Printf("[FileLogger| file: %s | level: %s] %s\n", logger.filePath, level, msg)
}

//
type dbLogger struct {
	database string
}

func (logger dbLogger) log(msg string, isError bool) {
	// fake implementation for demo
	var level string = "Info"
	if isError {
		level = "Error"
	}
	fmt.Printf("[DbLogger| database: %s | level: %s] %s\n", logger.database, level, msg)
}

//

func main() {
	fileLogger := fileLogger{filePath: "C:/logs.txt"}
	dbLogger := dbLogger{database: "system_logs"}

	fmt.Println("We can use interfaces to make our functions depend on abstraction instead of implementation (which is one way of achieving inversion of control)")
	errorMsg := "Something went wrong"
	logError(fileLogger, errorMsg)
	logError(dbLogger, errorMsg)

	infoMsg := "New user registered"
	logInfo(fileLogger, infoMsg)
	logInfo(dbLogger, infoMsg)

	fmt.Println("We can also use interfaces to create broader set of types matching our criteria for arrays, maps etc")
	loggers := []ilogger{fileLogger, dbLogger}

	for _, logger := range loggers {
		logger.log("Message from loop", false)
	}
}
