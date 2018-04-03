package main

import (
	"log"
	"regexp"

	"github.com/spf13/pflag"
)

var patternsToMatch []string
var regexpsToMatch []*regexp.Regexp
var screensaverDisableDelay int

func init() {
	pflag.StringArrayVar(&patternsToMatch, "match", []string{}, "Regexps to match processes on")
	pflag.IntVar(&screensaverDisableDelay, "delay", 30, "Disable delay in seconds")
	pflag.Parse()

	if err := initRegexps(); err != nil {
		log.Fatal(err)
	}

}

func initRegexps() (err error) {
	regexpsToMatch = make([]*regexp.Regexp, len(patternsToMatch))

	for i, pattern := range patternsToMatch {
		regexpsToMatch[i], err = regexp.Compile(pattern)

		if err != nil {
			return err
		}
	}

	return nil
}

func matchesAnyRegexp(input string) (matched bool) {
	for _, reg := range regexpsToMatch {
		matched = reg.MatchString(input)
		if matched {
			return matched
		}
	}

	return matched
}
