package commands

import (
	"fmt"

	"github.com/apereiroc/go-bp/internal/templates"
	"github.com/spf13/cobra"
)

func NewFileGeneratorCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "file [type] [output-path]",
		Short: "Generate a single-file template",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			templateType := args[0]
			outputPath := args[1]

			tmpl, err := templates.GetSingleFileTemplate(templateType)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				single := templates.ListTemplates()
				fmt.Printf("Supported single-file templates: %v\n", single)
				return
			}

			if err := tmpl.Generate(outputPath); err != nil {
				fmt.Printf("Error generating single template: %v\n", err)
			} else {
				fmt.Println("Template created successfully!")
			}
		},
	}
}
