package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Receipt struct {
	buf   string
	Width int
}

func (r *Receipt) Line(format string, params ...any) *Receipt {
	//s := fmt(fmt, params)
	s := fmt.Sprintf(format, params...)
	r.buf = r.buf + wrap(s, r.Width) + "\n"

	return r
}

func (r *Receipt) Rule() *Receipt {
	r.buf = r.buf + strings.Repeat("=", r.Width) + "\n"

	return r
}

func (r *Receipt) Divider(character string) *Receipt {
	r.buf = r.buf + strings.Repeat(character, r.Width) + "\n"

	return r
}

func (r *Receipt) Flush() string {
	return r.buf
}

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
