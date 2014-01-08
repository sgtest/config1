package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func CmdWaitWithTimeout(timeout time.Duration, cmd *exec.Cmd) error {
	errc := make(chan error, 1)
	go func() {
		errc <- cmd.Wait()
	}()
	var err error
	select {
	case <-time.After(timeout):
		cmd.Process.Kill()
		err = fmt.Errorf("timed out after %v", timeout)
	case err = <-errc:
	}
	return err
}

// ShellCommandString returns a shell command string that would run roughly the
// same command.
func ShellCommandString(cmd *exec.Cmd) string {
	parts := make([]string, 1+len(cmd.Args))
	parts[0] = cmd.Path
	for i, arg := range cmd.Args {
		escaped := fmt.Sprintf("%q", arg)
		parts[i+1] = escaped[1 : len(escaped)-1]
	}
	return strings.Join(parts, " ")
}
