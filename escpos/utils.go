package escpos

import "bytes"

func wrap(s string, length int) string {
	if len(s) == 0 {
		return ""
	}

	if length <= 0 {
		return ""
	}

	sub := ""

	runes := bytes.Runes([]byte(s))
	current := 0
	for _, r := range runes {
		current++
		c := string(r)

		if c == " " && current == 1 {
			current--
			continue
		}
		if c == "\n" || current == length {
			sub = sub + "\n"
			current = 0
			continue
		}

		sub = sub + c
	}

	return sub
}
