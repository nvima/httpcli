package util

import (
	"fmt"
	"os"
)

func HandleError(err error, msg string) {
    // fmt.Println(err)
	fmt.Println(msg)
	os.Exit(1)
}
