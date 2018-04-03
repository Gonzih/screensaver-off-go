package main

import (
	"os"
	"os/exec"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Gonzih/screensaver-off-go/icons"
	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/process"
)

type appState struct {
	manuallyDisabled      bool
	automaticallyDisabled bool
}

var (
	state                   appState
	stateLock               sync.RWMutex
	stateChangeNotification chan bool
)

func isAnyProcessMatching() bool {
	if len(regexpsToMatch) > 0 {
		procs, err := process.Processes()
		if err != nil {
			log.Infof("Error getting processes: %s", err)
			return false
		}

		var matched bool
		var name string

		for _, proc := range procs {
			name, err = proc.Name()
			if err != nil {
				log.Infof("Error getting process name: %s", err)
				continue
			}

			matched = matchesAnyRegexp(name)
			if matched {
				log.Infof(`Found a matching proc "%s"`, name)
				return matched
			}
		}
	}

	return false

}

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

func screensaverLoop() {
	log.Info("Starting screensaver loop")
	for {
		tryToDisable()
		time.Sleep(time.Second * time.Duration(screensaverDisableDelay))
	}
}

func toggleAutomaticState() {
	matched := isAnyProcessMatching()

	stateLock.Lock()
	defer stateLock.Unlock()

	if state.automaticallyDisabled != matched {
		state.automaticallyDisabled = matched
		stateChangeNotification <- true
	}
}

func procObserverLoop() {
	log.Info("Starting proc observer")
	for {
		toggleAutomaticState()
		time.Sleep(time.Second)
	}
}

func init() {
	stateChangeNotification = make(chan bool, 10)
}

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	go screensaverLoop()
	go procObserverLoop()

	systray.SetIcon(icons.IconEmpty)
	systray.SetTitle("Disable xscreensaver")
	systray.SetTooltip("Disable xscreensaver")
	mToggle := systray.AddMenuItem("Disable", "Disable")
	mQuit := systray.AddMenuItem("Quit", "Quit screensaver-off")

	go func() {
		for _ = range stateChangeNotification {
			log.Info("Got notified about state change")
			stateLock.RLock()

			if state.manuallyDisabled || state.automaticallyDisabled {
				mToggle.Check()
				mToggle.SetTitle("Enable")
				systray.SetIcon(icons.IconFull)
			} else {
				mToggle.Uncheck()
				mToggle.SetTitle("Disable")
				systray.SetIcon(icons.IconEmpty)
			}

			stateLock.RUnlock()
		}
	}()

	go func() {
		for {
			select {
			case <-mToggle.ClickedCh:
				stateLock.Lock()
				state.manuallyDisabled = !state.manuallyDisabled
				stateLock.Unlock()
				stateChangeNotification <- true
			case <-mQuit.ClickedCh:
				systray.Quit()
				break
			}
		}
	}()

	log.Info("Ready!")
}

func onExit() {
	os.Exit(0)
}
