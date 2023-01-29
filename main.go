package main

import (
	"log"
	"os"
	"sqler/cmdargs"

	slogger "github.com/jesse-rb/slogger-go"
)

var cmdArgs *cmdargs.CmdArgs

func init() {
    cmdArgs := cmdargs.NewArgs();
    cmdArgs.Expect("", "The sql file names to produce an ER diagram for")
    cmdArgs.Expect("nofkc", "Try to produce without considering foreign key constraints, only relying on column names match table names")
}

func main() {
    infoLogger := slogger.New(os.Stdout, slogger.ANSIBlue, "info", log.Lshortfile+log.Ldate);
    errorLogger := slogger.New(os.Stdout, slogger.ANSIRed, "error", log.Lshortfile+log.Ldate);
    infoLogger.Log("main", "Something worth noting happened", 2+4)
    errorLogger.Log("main", "Some horrible error happened", []int{3, 5, 7})
}
