package deployer

import "fmt"

type CMD struct {
	fmt.Stringer
	command string
	pid     int
}

func BuildCMD(cmd string) CMD {
	return CMD{command: cmd}
}

func (cmd *CMD) Execute() {
	pid := 0
	cmd.pid = pid
}

func (cmd *CMD) GetPid() int {
	return cmd.pid
}

func (cmd *CMD) GetCommand() string {
	return cmd.command
}

func (cmd *CMD) String() string {
	return fmt.Sprintf("CMD{command:%s, pid:%d}", cmd.GetCommand(), cmd.pid)
}
