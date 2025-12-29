package bebi

import (
	"strings"
	"unicode/utf8"
)

func MakeBubble(text string) string {
	lines := strings.Split(text, "\n")

	max := 0                     // max length
	for _, line := range lines { // finding max length line to determine border
		length := utf8.RuneCountInString(line)
		if length > max {
			max = length
		}
	}

	bubble := " "
	// top border
	bubble += strings.Repeat("_", max+2)
	bubble += "\n"

	for i, line := range lines {
		length := utf8.RuneCountInString(line)
		padding := strings.Repeat(" ", max-length) // add space for same width

		if len(lines) == 1 { // if only one line
			bubble += "< " + line + padding + " >\n"
		} else if i == 0 {
			bubble += "/ " + line + padding + " \\\n"
		} else if i == len(lines)-1 { // last line
			bubble += "\\ " + line + padding + " /\n"
		} else {
			bubble += "| " + line + padding + " |\n"
		}
	}

	// bottom border
	bubble += " " + strings.Repeat("-", max+2)
	return bubble
}
