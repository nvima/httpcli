package util

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func HandleError(cmd *cobra.Command, err error, tplError TplError) {
	// fmt.Println(err)
	fmt.Fprintf(cmd.OutOrStdout(), tplError.msg)
	os.Exit(tplError.exitCode)
}
