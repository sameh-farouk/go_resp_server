package commands

import (
	"errors"
	"net"
	"time"
)

func Hi() string {
	return "Hi back!"
}

func TestURL(args []string) (bool, error) {
	if len(args) < 1 {
		return false, errors.New("url arg is missing")
	}
	timeout := 2 * time.Second
	_, err := net.DialTimeout("tcp", args[0]+":http", timeout)
	if err != nil {
		return false, nil
	}
	return true, nil
}
