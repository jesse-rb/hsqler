package cmdargs

import (
	"fmt"
	"log"
	"os"

	logger "github.com/jesse-rb/slogger-go"
)

var infoLogger *logger.Logger

type CmdArgs struct {
    args map[string]*Arg
}

type Arg struct {
    name string
    short string
    desc string
    usageMsg string
    length int
    present bool
    value []string
}

func NewCmdArgs() (*CmdArgs) {
    infoLogger = logger.New(os.Stdout, logger.ANSIBlue, "[cmdargs]", log.Lshortfile+log.Ldate)

    var args map[string]*Arg = make(map[string]*Arg)
    var cmdArgs = &CmdArgs{ args: args }
    return cmdArgs
}

func newArg(name string, short string, desc string, usageMsg string, length int, defaultValue []string) (*Arg) {
    var arg = &Arg{name: name, short: short, desc: desc, usageMsg: usageMsg, length: length, present: false, value: defaultValue}
    return arg
}

func (ca *CmdArgs) Expect(name string, short string, desc string, usageMsg string, length int, defaultValue []string) {
    var newArg *Arg = newArg(name, short, desc, usageMsg, length, defaultValue)
    ca.args[name] = newArg
}

func (ca *CmdArgs) Parse(osArgs []string) {
    infoLogger.Log("Parse()", "Received os args array", osArgs)

    // If no os args, we are done
    if len(osArgs) == 0 {
        return;
    }

    var processed int = 1;

    // Check for the unnamed arg initially
    if _, ok := ca.args[osArgs[0]]; !ok {
        if theUnnamedArg, ok := ca.args[""]; ok {
            processed = ca.process(theUnnamedArg, osArgs, processed)
        }
    }
    // Now check for any other named args afterwards
    for _, a := range osArgs {
        if theArg, ok := ca.args[a]; ok {
            processed = ca.process(theArg, osArgs, processed);
        }
    }
    
    infoLogger.Log("[parse cmd args]", "finished", ca.ToString())
}

// Get an arg by a key string
// returns nil if arg not found
func (ca *CmdArgs) Get(key string) (*Arg) {
    if arg, ok := ca.args[key]; ok {
        return arg
    }
    return nil
}

func (a *Arg) Present() (bool) {
    return a.present
}
func (a *Arg) Length() (int) {
    return a.length;
}
func (a *Arg) GetValue() ([]string) {
    return a.value
}

func (ca *CmdArgs) process(arg *Arg, osArgs []string, processed int) (int) {
    // Arg is present
    arg.present = true
    if (arg.length == -1) { // Arg value has infinite capacity
        for processed < len(osArgs) {
            if _, ok := ca.args[osArgs[processed]]; ok {
                break
            }
            arg.value = append(arg.value, osArgs[processed])
            processed++
        }
    } else { // Arg value has finite capacity
        for i := 0; i < arg.length; i++ {
            arg.value = append(arg.value, osArgs[processed])
            processed++
        }
    }
    return processed
}

func (ca *CmdArgs) ToString() (string) {
    var s string = "args: [\n"
    for i, a := range ca.args {
        s += fmt.Sprintf("\t%v: %v,\n", i, a.ToString())
    }
    s += "]"
    return s
}

func (a *Arg) ToString() (string) {
    var s string = "{\n";
    s += fmt.Sprintf("\t\tname: %v,\n\t\tshort: %v,\n\t\tdesc: %v,\n\t\tusageMsg: %v,\n\t\tpresent: %v,\n\t\t,\n\t\tlength: %v\n", 
        a.name, a.short, a.desc, a.usageMsg, a.present, a.length)
    s += "\t\tvalue: [\n"
    for i, v := range a.value {
        s += fmt.Sprintf("\t\t\t%v: %v\n", i, v)
    }
    s += "\t\t]\n"
    s += "\t}"
    return s
}

