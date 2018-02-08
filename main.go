package main

import (
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/Gonzih/screensaver-off-go/icons"
	"github.com/getlantern/systray"
)

var (
	disableScreensaver     bool
	disableScreensaverLock sync.RWMutex
)

func tryToDisable() {
	disableScreensaverLock.RLock()
	defer disableScreensaverLock.RUnlock()

	if disableScreensaver {
		log.Println("Trying to disable xscreensaver")
		cmd := exec.Command("/usr/bin/xscreensaver-command", "-deactivate")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Printf("Error executing xscreensaver-command: %s\n", err)
		}
	}
}

func screensaverLoop() {
	for {
		tryToDisable()
		time.Sleep(time.Second)
	}
}

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icons.IconFull)
	systray.SetTitle("Disable xscreensaver")
	systray.SetTooltip("Disable xscreensaver")
	mToggle := systray.AddMenuItem("Disable", "Disable")
	mQuit := systray.AddMenuItem("Quit", "Quit screensaver-off")

	go func() {
		for {
			select {
			case <-mToggle.ClickedCh:

				disableScreensaverLock.Lock()
				disableScreensaver = !disableScreensaver
				disableScreensaverLock.Unlock()

				disableScreensaverLock.RLock()

				if disableScreensaver {
					mToggle.Check()
					mToggle.SetTitle("Enable")
					systray.SetIcon(icons.IconEmpty)
				} else {
					mToggle.Uncheck()
					mToggle.SetTitle("Disable")
					systray.SetIcon(icons.IconFull)
				}

				disableScreensaverLock.RUnlock()

			case <-mQuit.ClickedCh:
				systray.Quit()
				break
			}
		}
	}()

	log.Println("Ready!")
}

func onExit() {
	os.Exit(0)
}
