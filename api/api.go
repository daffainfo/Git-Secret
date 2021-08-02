package api

import (
	"fmt"
	"regexp"
)

func google_api(contents string) {
	re := regexp.MustCompile(`AIza[0-9A-Za-z-_]{35}`)
	fmt.Println("Google API Key:", re.MatchString(contents))
}

func twitter_secret(contents string) {
	re := regexp.MustCompile(`(?i)twitter(.{0,20})?[0-9a-z]{35,44}`)
	fmt.Println("Twitter Secret:", re.MatchString(contents))
}

func twilio_api(contents string) {
	re := regexp.MustCompile(`(?i)twilio(.{0,20})?SK[0-9a-f]{32}`)
	fmt.Println("Twilio API:", re.MatchString(contents))
}

func stripe(contents string) {
	re := regexp.MustCompile(`(?i)twilio(.{0,20})?SK[0-9a-f]{32}`)
	fmt.Println("Twilio API:", re.MatchString(contents))
}
