package main

import (
	"github.com/shirou/gopsutil/process"
	log "github.com/sirupsen/logrus"
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
