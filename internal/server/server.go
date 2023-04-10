package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/sameh-farouk/go_resp_server/internal/commands"
	"github.com/sameh-farouk/go_resp_server/pkg/resp"
)

func StartServer(host string, port string, _type string) {
	l, err := net.Listen(_type, host+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + host + ":" + port)
	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	fmt.Println("Accepted")

	r := bufio.NewReader(conn)
	respReader := resp.NewReader(r)
	for {
		// detect if the client close the connection, no further reading required
		_, err := r.Peek(1)
		if err == io.EOF {
			break
		}
		command, args, err := respReader.ParseCommand()
		if err != nil {
			fmt.Println("Error: ", err)
			break

		}
		fmt.Println("command: " + command)
		fmt.Println("args: ", args)

		switch command {
		case "hi":
			res := commands.Hi()
			conn.Write([]byte("$" + fmt.Sprint(len(res)) + "\r\n" + res + "\r\n"))

		case "testurl":
			res, err := commands.TestURL(args)
			if err != nil {
				conn.Write([]byte("-" + err.Error() + "\r\n"))
				break
			}
			var bitSetVar int8
			if res {
				bitSetVar = 1
			}
			conn.Write([]byte(":" + fmt.Sprint(bitSetVar) + "\r\n"))
		default:
			conn.Write([]byte("-unknown command\r\n"))
		}
	}
	conn.Close()
	fmt.Println("Connection Closed")
}
