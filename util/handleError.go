package util

import (
	"fmt"

	"github.com/spf13/cobra"
)

func HandleError(cmd *cobra.Command, err error, tplError error) error {
	debug, _ := cmd.Flags().GetBool("debug")
	if debug {
		fmt.Printf("Debug: %v\n", err)
	}
	fmt.Println(tplError.Error())
	return tplError
}
