package main

import (
	"log"
	"os"
	"sqler/cmdargs"
	"sqler/enums/logtags"
	"sqler/sqlparser"

	slogger "github.com/jesse-rb/slogger-go"
)

var cmdArgs *cmdargs.CmdArgs

func init() {
    cmdArgs := cmdargs.NewCmdArgs()
    cmdArgs.Expect("", "", "The sql file name(s) to produce an ER diagram for", "<file name 1> <file name 2> ... <file name n>", -1, nil)
    cmdArgs.Expect("--nofkc", "-n", "Try to produce without considering foreign key constraints, only relying on column names that match table names", "--nofkc", 0, nil)
    cmdArgs.Parse(os.Args)
}

func main() {
    infoLogger := slogger.New(os.Stdout, slogger.ANSIBlue, "info", log.Ldate);
    errorLogger := slogger.New(os.Stdout, slogger.ANSIRed, "error", log.Ldate);

    infoLogger.Log(logtags.Main, "init", 1)
    errorLogger.Log(logtags.Main, "init", 1)

    // Check input
    files := make([]string, 0)

    if cmdArgs.Get("") != nil {
        var argFiles *cmdargs.Arg = cmdArgs.Get("")
        files = argFiles.GetValue();
    } 

    sqlparser.Parse(files);
}
