package main

import (
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
    slogger.LogInfo("main", "Some info.", 2+4)
    slogger.LogError("main", "Some error", 4+0)
}
