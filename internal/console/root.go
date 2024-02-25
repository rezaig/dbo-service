package console

import (
	"fmt"
	"os"

	"github.com/rezaig/dbo-service/internal/config"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: "dbo-service",
}

func init() {
	config.GetConfig()
}

// Execute handles command execution
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
