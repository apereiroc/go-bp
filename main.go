package main

import (
	"fmt"
	"os"

	"github.com/apereiroc/go-bp/internal/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: os.Args[0]}

	rootCmd.AddCommand(commands.NewFileGeneratorCmd())
	rootCmd.AddCommand(commands.NewListTemplatesCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
