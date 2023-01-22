package main

import (
    "sqler/logger"
)

func init() {

}

func main() {
    logger.LogError("main", "something horrible happened.", []int{1,2,3})
    logger.LogInfo("main", "Something worth noting happened.", []int{1,2,3})

}
