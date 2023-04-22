package util

import (
	// "fmt"

	"github.com/spf13/cobra"
)

func HandleError(cmd *cobra.Command, err error, tplError TplError) error {
	//TODO:: Add Error Logging with Stacktraces
	// fmt.Println(err)

	// fmt.Fprintf(cmd.OutOrStdout(), tplError.msg)
	return err
}
