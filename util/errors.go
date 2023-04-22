package util

import (
	"fmt"
)

type TplError struct {
	msg      string
	err      error
	exitCode int
}

func (te TplError) Err() error {
	return te.err
}

func (te TplError) Msg() string {
	return te.msg
}

func (te TplError) ExitCode() int {
	return te.exitCode
}

var NO_CONFIG_FILE_ERR = TplError{
	msg:      "No config file found",
	err:      fmt.Errorf("No config file found"),
	exitCode: 1,
}

var NO_FUNC_NAME_ERR = TplError{
	msg:      "No function name provided",
	err:      fmt.Errorf("No function name provided"),
	exitCode: 2,
}
var NO_FUNC_FOUND_ERR = TplError{
    msg:      "Function not found in config",
    err:      fmt.Errorf("Function not found in config"),
    exitCode: 3,
}

var INVALID_CONFIG_ERR = TplError{
    msg:      "Invalid config file",
    exitCode: 4,
}

var INVALID_RESP_CODE = TplError{
    msg:      "Invalid response code",
    exitCode: 5,
}
var FAILED_TO_GET_DATA = TplError{
    msg:      "Failed to get data",
    exitCode: 6,
}
var FAILED_TO_MAKE_HTTP_CALL = TplError{
    msg:      "Failed to make http call",
    exitCode: 7,
}
var FAILED_TO_PARSE_JSON = TplError{
    msg:      "Failed to parse JSON",
    exitCode: 8,
}
var FAILED_TO_PARSE_OUTPUT_FIELD = TplError{
    msg:      "Failed to parse output field",
    exitCode: 9,
}
