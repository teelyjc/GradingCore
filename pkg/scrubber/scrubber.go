package scrubber

import (
	"strings"
)

func Scrub(data []byte, isAutoTrimEnabled bool) []byte {
	return []byte(ScrubString(string(data), isAutoTrimEnabled))
}

func ScrubString(data string, isAutoTrimEnabled bool) string {
	split := strings.Split(data, "\n")
	splitlen := len(split)
	for i := 0; i < splitlen; i++ {
		split[i] = strings.TrimSuffix(split[i], "\r")
		if isAutoTrimEnabled {
			split[i] = strings.TrimSpace(split[i])
		}
	}

	if split[splitlen-1] == "" {
		split = split[0 : splitlen-1]
	}

	return strings.Join(split, "\n") + "\n"
}
