// +build linux

package url

import (
	"os/exec"
)

func openUrl(url string) {
	exec.Command("/usr/bin/env", "xdg-open", url).Run()
	// Support for WSL
	exec.Command("/usr/bin/env", "cmd.exe", "/C", "start", url).Run()
}
