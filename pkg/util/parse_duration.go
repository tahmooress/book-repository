package util

import "time"

// ParseDurationWithDefault get a duration in string format and a default value
// and parse the string in time.Duration type and if it fale return
// the default value.
func ParseDurationWithDefault(strDuration string, deaultValue time.Duration) time.Duration {
	t, err := time.ParseDuration(strDuration)
	if err != nil {
		t = deaultValue
	}

	return t
}
