// +build freebsd

package url

import (
	"os/exec"
)

func openUrl(url string) {
	exec.Command("/usr/local/bin/xdg-open", url).Run()
}
