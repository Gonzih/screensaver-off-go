package main

import (
	"flag"
	"log"
	"regexp"
)

type stringArrayFlag []string

func (i *stringArrayFlag) String() string {
	return "my string representation"
}

func (i *stringArrayFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var patternsToMatch stringArrayFlag
var regexpsToMatch []*regexp.Regexp
var screensaverDisableDelay int

func init() {
	initFlags()

	if err := initRegexps(); err != nil {
		log.Fatal(err)
	}

}

func initFlags() {
	flag.Var(&patternsToMatch, "match", "Regexp to match processes on")
	flag.IntVar(&screensaverDisableDelay, "delay", 30, "Disable delay in seconds")
	flag.Parse()
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
