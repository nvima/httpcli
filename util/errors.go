package util

import (
	"errors"
)

var NO_CONFIG_FILE_ERR = errors.New("No config file found")
var NO_FUNC_NAME_ERR = errors.New("No function name provided")
var NO_FUNC_FOUND_ERR = errors.New("Function not found in config")
var INVALID_CONFIG_ERR = errors.New("Invalid config file")
var INVALID_RESP_CODE = errors.New("Invalid response code")
var FAILED_TO_GET_DATA = errors.New("Failed to get data")
var FAILED_TO_MAKE_HTTP_CALL = errors.New("Failed to make http call")
var FAILED_TO_PARSE_JSON_RESPONSE = errors.New("Failed to parse JSON Response")
var FAILED_TO_PARSE_OUTPUT_FIELD = errors.New("Failed to parse output field")
var NO_URL_ERR = errors.New("No URL provided")
var MARSHAL_DATA_FAILED = errors.New("Failed to marshal data")
var REPLACE_STDIN_FAILED = errors.New("Failed to replace stdin")
var INIT_HTTP_POST_REQUEST_FAILED = errors.New("An error occurred while initializing the HTTP POST request. Possible causes could be an invalid URL or incorrect input parameters.")
var SEND_HTTP_POST_REQUEST_FAILED = errors.New("An error occurred while sending the HTTP POST request. Possible causes could be network issues, server unavailability, or issues with the request payload.")
var READ_RESPONSE_BODY_FAILED = errors.New("Failed to read the response body")
