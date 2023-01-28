package cmdargs

type CmdArgs struct {
    Args map[string][]string
}

func NewArgs() (*CmdArgs) {
    var args map[string][]string = make(map[string][]string)
    var cmdArgs = &CmdArgs{ Args: args }
    return cmdArgs
}

func (ca *CmdArgs) Expect(name string, desc string) {
    
}

func Parse(in []string) {

}

func ParseUnnamed(in string) {

}

func ParseNamed(in string) {

}
