package commands

import (
	"fmt"

	"github.com/apereiroc/go-bp/internal/templates"
	"github.com/spf13/cobra"
)

func NewListTemplatesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Lists all available templates",
		Run: func(cmd *cobra.Command, args []string) {
			single := templates.ListTemplates()
			fmt.Println("Available single-file templates:")
			for _, t := range single {
				fmt.Printf("  - %s\n", t)
			}
		},
	}
}
