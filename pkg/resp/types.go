package resp

type RespType int

const (
	SIMPLE_STRING_PREFIX = byte('+')
	ERROR_PREFIX         = byte('-')
	INTEGER_PREFIX       = byte(':')
	BULK_STRING_PREFIX   = byte('$')
	ARRAY_PREFIX         = byte('*')
)

const (
	SimpleString RespType = iota
	Error
	Integer
	BulkString
	Array
)

type RespData struct {
	Data []string
	Type RespType
}

func NewRespData(data []byte) {

}
