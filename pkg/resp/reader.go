package resp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type RESPReader struct {
	*bufio.Reader
}

func NewReader(reader io.Reader) *RESPReader {
	return &RESPReader{
		Reader: bufio.NewReaderSize(reader, 32*1024),
	}
}

func (r *RESPReader) ParseCommand() (string, []string, error) {
	// validate type
	firstToken, err := r.ReadBytes(byte('\n'))
	if err != nil && err != io.EOF {
		return "", nil, errors.New("not array! can't parse the command")
	}
	if err == io.EOF {
		fmt.Println("EOF")
	}
	i64, _ := strconv.ParseInt(string(firstToken[1:len(firstToken)-2]), 10, 0)
	_len := int(i64)

	if _len < 1 {
		return "", nil, errors.New("len < 1! can't parse the command")
	}

	cmd, err := r.ReadBulkString()
	if err != nil {
		return "", nil, err
	}
	var args []string
	for i := 1; i < _len; i++ {
		arg, err := r.ReadBulkString()
		if err != nil {
			return "", nil, err
		}
		args = append(args, arg)
	}

	return cmd, args, nil

}

func (r *RESPReader) ReadBulkString() (string, error) {
	firstToken, err := r.ReadBytes(byte('\n'))
	if err != nil {
		return "", errors.New("can't find delimiter! can't parse the bulk string\n" + err.Error())
	}
	if firstToken[0] != BULK_STRING_PREFIX {
		fmt.Println(firstToken)
		return "", errors.New("not bulk string! can't parse the bulk string ")
	}
	i64, _ := strconv.ParseInt(string(firstToken[1:len(firstToken)-2]), 10, 0)
	_len := int(i64)
	// TODO: handle empty and null
	buf := make([]byte, _len)
	n, _ := io.ReadFull(r, buf)
	if n != _len {
		return "", errors.New("can't parse the bulk string")
	}
	r.Discard(2)
	return string(buf), nil
}
