package main

import (
	"fmt"
	"sqler/logger"
)

func init() {

}

func main() {
    fmt.Println("test")
    var test string = "hello";
    test = test + " " + "world"
    logger.LogInfo("main", "this is a test log", test)
}
