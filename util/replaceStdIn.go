package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func removeControlChars(s string) string {
	return strings.Map(func(r rune) rune {
		if r == '\t' {
			return ' '
		}
		if r == '\n' || unicode.IsPrint(r) {
			return r
		}
		return -1
	}, s)
}

func ReplaceStdIn(input []byte) ([]byte, error) {
	inputStr := string(input)
	if strings.Contains(inputStr, "${STDIN}") {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("Error reading stdin: %v", err)
		}

		stdInStr := string(data)
		stdInStr = removeControlChars(stdInStr)

		escapedStdIn, err := json.Marshal(stdInStr)
		if err != nil {
			return nil, fmt.Errorf("Error escaping stdin string: %v", err)
		}
		stdInStr = string(escapedStdIn[1 : len(escapedStdIn)-1])

		result := strings.Replace(inputStr, "${STDIN}", stdInStr, -1)
		return []byte(result), nil
	}
	return input, nil
}

func ParseJSONResponse(jsonData []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetOutputField(data interface{}, fieldPath string) string {
	keys := strings.Split(fieldPath, ".")

	var result interface{} = data
	for _, key := range keys {
		if key == "" {
			continue
		}
		if strings.Contains(key, "[") {
			innerKey := key[:strings.Index(key, "[")]
			index := key[strings.Index(key, "[")+1 : strings.Index(key, "]")]
			m, ok := result.(map[string]interface{})[innerKey].([]interface{})
			if !ok {
				panic("invalid output path")
			}
			intVar, _ := strconv.Atoi(index)
			result = m[intVar]
		} else {
			m, ok := result.(map[string]interface{})[key]
			if !ok {
				panic("invalid output path")
			}
			result = m
		}
	}

	if _, ok := result.(map[string]interface{}); ok {
		jsonResult, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}
		return string(jsonResult)
	}
	if _, ok := result.(string); ok {
		return result.(string)
	}
	panic("invalid output path")
}
