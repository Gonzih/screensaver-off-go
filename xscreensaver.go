package main

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func tryToDisable() {
	stateLock.RLock()
	defer stateLock.RUnlock()

	if state.manuallyDisabled || state.automaticallyDisabled {
		log.Info("Trying to disable xscreensaver")
		cmd := exec.Command("/usr/bin/xscreensaver-command", "-deactivate")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Errorf("Error executing xscreensaver-command: %s", err)
		}
	}
}
