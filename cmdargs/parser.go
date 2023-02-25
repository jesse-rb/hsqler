package cmdargs

import (
	"log"
	"os"

	logger "github.com/jesse-rb/slogger-go"
)

var infoLogger *logger.Logger

type CmdArgs struct {
    args map[string]*arg
}

type arg struct {
    name string
    short string
    desc string
    usageMsg string
    length int
    present bool
    value *[]string
}

func NewArgs() (*CmdArgs) {
    infoLogger = logger.New(os.Stdout, logger.ANSIBlue, "[cmdargs]", log.Lshortfile+log.Ldate)

    var args map[string]*arg = make(map[string]*arg)
    var cmdArgs = &CmdArgs{ args: args }
    return cmdArgs
}

func newArg(name string, short string, desc string, usageMsg string, length int, defaultValue *[]string) (*arg) {
    var arg = &arg{name: name, short: short, desc: desc, usageMsg: usageMsg, length: length, present: false, value: defaultValue}
    return arg
}

func (ca *CmdArgs) Expect(name string, short string, desc string, usageMsg string, length int, defaultValue *[]string) {
    var newArg *arg = newArg(name, short, desc, usageMsg, length, defaultValue)
    ca.args[name] = newArg
}

func (ca *CmdArgs) Parse(osArgs []string) {
    infoLogger.Log("Parse()", "Received os args array", osArgs)

    // If no os args, we are done
    if len(osArgs) == 0 {
        return;
    }

    // Check for the unnamed arg initially
    if arg, ok := ca.args[osArgs[0]]; !ok {
        infoLogger.Log("Parse()", "Trying to get arg", arg)
    }
    // Now check for any other named args afterwards
    for _, a := range osArgs {
        infoLogger.Log("Parse()", "Looping over os args", a)
        if arg, ok := ca.args[a]; ok {
            infoLogger.Log("Parse()", "Trying to get arg", arg)
        }
    }
}
