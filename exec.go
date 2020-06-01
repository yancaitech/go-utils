package utils

import (
	"os/exec"
)

// Exec command and output STDOUT
func Exec(cmd string) (out string, err error) {
	ob, err := exec.Command(cmd).Output()
	if err != nil {
		out = string(ob)
	}
	return out, err
}
