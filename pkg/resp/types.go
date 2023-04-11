package resp

const (
	SIMPLE_STRING_PREFIX = byte('+')
	ERROR_PREFIX         = byte('-')
	INTEGER_PREFIX       = byte(':')
	BULK_STRING_PREFIX   = byte('$')
	ARRAY_PREFIX         = byte('*')
)
