package main

import (
	"github.com/sameh-farouk/go_resp_server/internal/server"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	server.StartServer(CONN_HOST, CONN_PORT, CONN_TYPE)
}
