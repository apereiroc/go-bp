package main

import (
	"fmt"
	"os"

	"github.com/apereiroc/temple/internal/commands"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: os.Args[0]}

	rootCmd.AddCommand(commands.NewFileGeneratorCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
