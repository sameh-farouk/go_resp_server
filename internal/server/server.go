package server

import (
	"bufio"
	"io"
	"log"
	"net"

	"github.com/sameh-farouk/go_resp_server/internal/commands"
	"github.com/sameh-farouk/go_resp_server/pkg/resp"
)

type RespServer struct {
	host    string
	port    string
	network string
}

func New(host string, port string, network string) RespServer {
	return RespServer{
		host,
		port,
		network,
	}
}

func (s *RespServer) Listen() {
	l, err := net.Listen(s.network, s.host+":"+s.port)
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}

	defer l.Close()
	log.Println("Listening on " + s.host + ":" + s.port)
	for {

		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		go s.handleRequest(conn)
	}
}

// Handles incoming requests.
func (s *RespServer) handleRequest(conn net.Conn) {
	log.Println("Accepted new connection")

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
			log.Println("Error: ", err)
			break

		}
		log.Println("command: ", command, "args: ", args)

		// TODO: move to functions on septate handler type
		switch command {
		case "HI":

			res := commands.Hi()
			respLine := resp.NewRespBulkString(res)
			conn.Write(respLine)

		case "TESTURL":
			res, err := commands.TestURL(args)
			if err != nil {
				conn.Write(resp.NewRespError(err.Error()))
				break
			}
			var bitSetVar int64
			if res {
				bitSetVar = 1
			}
			respLine := resp.NewRespInteger(bitSetVar)
			conn.Write(respLine)
		default:
			conn.Write(resp.NewRespError("unknown command"))
		}
	}
	conn.Close()
	log.Println("Connection Closed")
}
