package resp

import (
	"log"
	"strings"
)

func NewError(msg string) string {
	if strings.ContainsAny(msg, "\n\r") {
		log.Println("no newlines are allowed")

	}
	return "-" + msg + "\r\n"
}
