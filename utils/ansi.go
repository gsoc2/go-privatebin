package utils

import (
	"regexp"
)

var regex = regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")

// StripANSI strips ansi out of a string.
func StripANSI(str string) string {
	return regex.ReplaceAllString(str, "")
}
