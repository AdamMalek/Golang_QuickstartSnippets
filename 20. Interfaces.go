package main

import "fmt"

// Structs containing function with same signature
type fileLogger struct {
	filePath string
}

func (logger *fileLogger) log(msg string, isError bool) {
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

func (logger *dbLogger) log(msg string, isError bool) {
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

	errorMsg := "Something went wrong"
	logError(&fileLogger, errorMsg) // we need to use & as function logError uses pointer receiver
	logError(&dbLogger, errorMsg)

	infoMsg := "New user registered"
	logInfo(&fileLogger, infoMsg)
	logInfo(&dbLogger, infoMsg)

}

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
