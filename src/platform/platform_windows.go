//go:build windows

package platform

import (
	"errors"
	"fmt"
	"os/exec"
)

func GetSongData() (string, error) {
	out, err := exec.Command("cscript.exe", "//Nologo", "platform\\windows.js").Output()

	if err != nil {
		errors.New(fmt.Sprintf("%s\n", err))
	}

	return string(out), nil
}
