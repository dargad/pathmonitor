package pathmonitor

import (
	"errors"
	"os/exec"
	"strings"
)

const (
	WSSEP           = " "
	PLACEHOLDER     = "{}"
	ESC_PLACEHOLDER = "'{}'"
)

func ExecuteCommand(command string) (bool, error) {
	Trace.Println("Executing command", command)
	args := strings.Split(command, WSSEP)
	var cmd *exec.Cmd
	l := len(args)
	switch {
	case l > 1:
		cmd = exec.Command(args[0], args[1:]...)
	case l == 1:
		cmd = exec.Command(command)
	default:
		return false, errors.New("No or unknown command to execute.")
	}
	err := cmd.Run()
	return err == nil, err
}

func ReplacePlaceholders(command string, filename string) string {
	r := strings.NewReplacer(ESC_PLACEHOLDER, ESC_PLACEHOLDER, PLACEHOLDER, filename)
	return r.Replace(command)
}
