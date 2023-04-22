package util

import (
	"fmt"
)

type TplError struct {
	msg      string
	err      error
}

func (te TplError) Err() error {
	return fmt.Errorf(te.msg)
}

func (te TplError) Msg() string {
	return te.msg
}

var NO_CONFIG_FILE_ERR = TplError{
	msg:      "No config file found",
}

var NO_FUNC_NAME_ERR = TplError{
	msg:      "No function name provided",
}
var NO_FUNC_FOUND_ERR = TplError{
	msg:      "Function not found in config",
}

var INVALID_CONFIG_ERR = TplError{
	msg:      "Invalid config file",
}

var INVALID_RESP_CODE = TplError{
	msg:      "Invalid response code",
}
var FAILED_TO_GET_DATA = TplError{
	msg:      "Failed to get data",
}
var FAILED_TO_MAKE_HTTP_CALL = TplError{
	msg:      "Failed to make http call",
}
var FAILED_TO_PARSE_JSON = TplError{
	msg:      "Failed to parse JSON",
}
var FAILED_TO_PARSE_OUTPUT_FIELD = TplError{
	msg:      "Failed to parse output field",
}
var NO_URL_ERR = TplError{
	msg:      "No URL provided",
}
