package scrubber

import (
	"strings"
)

func Scrub(data []byte, autoTrimEnabled bool) []byte {
	return []byte(ScrubString(string(data), autoTrimEnabled))
}

func ScrubString(data string, autoTrimEnabled bool) string {
	split := strings.Split(data, "\n")
	splitlen := len(split)
	for i := 0; i < splitlen; i++ {
		split[i] = strings.TrimSuffix(split[i], "\r")
		if autoTrimEnabled {
			split[i] = strings.TrimSpace(split[i])
		}
	}

	if split[splitlen-1] == "" {
		split = split[0 : splitlen-1]
	}

	return strings.Join(split, "\n") + "\n"
}
