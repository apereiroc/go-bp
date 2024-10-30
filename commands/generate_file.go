package commands

import (
	"fmt"
	"strings"

	"github.com/apereiroc/go-bp/templates"
	"github.com/spf13/cobra"
)

func NewFileGeneratorCmd() *cobra.Command {
	// Fetch available single-file templates at command definition
	single := templates.ListTemplates()
	availableTemplates := "No single-file templates are currently available."
	if len(single) > 0 {
		availableTemplates = "Available template types:\n- " + strings.Join(single, "\n- ")
	}

	return &cobra.Command{
		Use:   "file [type] [output-path]",
		Short: "Generate a single-file template",
		Long:  fmt.Sprintf("Generates a single template file.\n\n%s", availableTemplates),
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			templateType := args[0]
			outputPath := args[1]

			tmpl, err := templates.GetSingleFileTemplate(templateType)
			if err != nil {
				return err
			}

			return tmpl.Generate(outputPath)
		},
	}
}
