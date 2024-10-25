package commands

import (
	"fmt"
	"log"

	"github.com/apereiroc/temple/internal/templates"
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

			var tmpl templates.Template

			switch templateType {
			case "makefile":
				// Read user input
				var author, projectName string

				fmt.Print("Enter author: ")
				_, err := fmt.Scanln(&author)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Print("Enter project name: ")
				_, err = fmt.Scanln(&projectName)
				if err != nil {
					log.Fatal(err)
				}

				tmpl = templates.MakefileTemplate{Author: author, ProjectName: projectName}

			default:
				fmt.Printf("Unknown template type: %s\n", templateType)
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
