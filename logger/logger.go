package logger

import (
    "log"
    "fmt"
    "os"
)

// 8 bit color ansi color codes for colored logging output
var ansiColors map[string]string = map[string]string {
    "reset": "\u001b[0m",
    "red": "\u001b[31m",
    "blue": "\u001b[34m",
}

// Init logger prefix variables
var globalPrefix = "sqler -> "
var errorPrefix = fmt.Sprintf("%serror -> ", ansiColors["red"])
var infoPrefix = fmt.Sprintf("%sinfo -> ", ansiColors["blue"])

// Init logger variables
var errorLogger *log.Logger = log.New(os.Stderr, fmt.Sprintf("%s%s", globalPrefix, errorPrefix), log.Flags())
var infoLogger *log.Logger = log.New(os.Stdout, fmt.Sprintf("%s%s", globalPrefix, infoPrefix), log.Flags())

// Internal log function that exported functions can call
func logGeneral(log *log.Logger, tag string, msg string, data interface{}) {
    log.Println(fmt.Sprintf("%s -> %s -> data:\n%d%s", tag, msg, data, ansiColors["reset"]));
}

// Log an error using the error logger
// example usage: logger.LogError("main", "something horrible happened.", []int{1,2,3})
func LogError(tag string, msg string, data interface{}) {
    logGeneral(errorLogger, tag, msg, data)
}

// Log info using the info logger
// example usage: logger.LogInfo("main", "Something worth noting happened.", []int{1,2,3})
func LogInfo(tag string, msg string, data interface{}) {
    logGeneral(infoLogger, tag, msg, data)
}
