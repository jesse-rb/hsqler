package logger

import (
    "log"
    "fmt"
    "os"
)

var ansiColors map[string]string = map[string]string {
    "reset": "\u001b[0m",
    "red": "\u001b[31m",
    "blue": "\u001b[34m",
}

var globalPrefix = "sqler -> "
var errorPrefix = fmt.Sprintf("%serror -> ", ansiColors["red"])
var infoPrefix = fmt.Sprintf("%sinfo -> ", ansiColors["blue"])

var errorLogger *log.Logger = log.New(os.Stderr, fmt.Sprintf("%s%s", globalPrefix, errorPrefix), log.Flags())
var infoLogger *log.Logger = log.New(os.Stdout, fmt.Sprintf("%s%s", globalPrefix, infoPrefix), log.Flags())

func logGeneral(log *log.Logger, tag string, msg string, data interface{}) {
    log.Println(fmt.Sprintf("%s -> %s -> data:\n%d%s", tag, msg, data, ansiColors["reset"]));

}

func LogError(tag string, msg string, data interface{}) {
    logGeneral(errorLogger, tag, msg, data)
}

func LogInfo(tag string, msg string, data interface{}) {
    logGeneral(infoLogger, tag, msg, data)
}
