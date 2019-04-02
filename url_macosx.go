// +build darwin

package url

import (
	"os"
	"os/exec"
	"log"
)

func openUrl(url string) {
	browser := ""

	if _, err := os.Stat("/Applications/Google Chrome.app"); err == nil {
		// path/to/whatever exists
		browser = "Google Chrome"
	} else if _, err := os.Stat("/Applications/Firefox.app"); err == nil {
		browser = "Firefox"
	} else if _, err := os.Stat("/Applications/Safari.app"); err == nil {
		browser = "Safari"
 	} else {
		log.Fatalf("Error: Chrome, Firefox or Safari not found, exiting")
	}

	log.Printf("Found browser: %s", browser)
	exec.Command("/usr/bin/open", "-a", browser, url).Run()
}
