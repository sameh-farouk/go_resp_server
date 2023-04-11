package resp

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

const (
	READER_INITIAL_SIZE = 32 * 1024
)

type RESPReader struct {
	*bufio.Reader
}

func NewReader(reader io.Reader) *RESPReader {
	return &RESPReader{
		Reader: bufio.NewReaderSize(reader, READER_INITIAL_SIZE),
	}
}

// this function expects only a RESP Array consisting of only Bulk Strings
// this is how redis sends commands to the server
func (r *RESPReader) ParseCommand() (string, []string, error) {
	// validate type
	firstToken, err := r.ReadBytes(byte('\n'))
	if err != nil {
		return "", nil, errors.New("can't find delimiter! can't parse the array\n" + err.Error())
	}
	if firstToken[0] != ARRAY_PREFIX {
		return "", nil, errors.New("not array! can't parse the bulk string ")
	}
	i64, _ := strconv.ParseInt(string(firstToken[1:len(firstToken)-2]), 10, 0)
	_len := int(i64)

	if _len < 1 {
		return "", nil, errors.New("len < 1! can't parse the command " + string(firstToken))
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

// parse a bulk string from the reader and advance it to next token
func (r *RESPReader) ReadBulkString() (string, error) {
	firstToken, err := r.ReadBytes(byte('\n'))
	if err != nil {
		return "", errors.New("can't find delimiter! can't parse the bulk string\n" + err.Error())
	}
	if firstToken[0] != BULK_STRING_PREFIX {
		return "", errors.New("not bulk string! can't parse the bulk string ")
	}
	i64, _ := strconv.ParseInt(string(firstToken[1:len(firstToken)-2]), 10, 0)
	_len := int(i64)
	// TODO: handle empty 0 and null -1
	buf := make([]byte, _len)
	n, _ := io.ReadFull(r, buf)
	if n != _len {
		return "", errors.New("can't parse the bulk string")
	}
	r.Discard(2)
	return string(buf), nil
}
