package resp

import (
	"fmt"
	"strings"
)

func NewRespError(err string) []byte {
	err = strings.ReplaceAll(err, "\n\r", "  ")
	return []byte("-" + err + "\r\n")
}

func NewSimpleString(str string) []byte {
	str = strings.ReplaceAll(str, "\n\r", "  ")
	return []byte("-" + str + "\r\n")
}

func NewRespInteger(i int64) []byte {
	return []byte(":" + fmt.Sprint(i) + "\r\n")
}

func NewRespBulkString(str string) []byte {
	bytes := []byte(str)
	return []byte("$" + fmt.Sprint(len(bytes)) + "\r\n" + str + "\r\n")
}

// TODO: add resp array writer
