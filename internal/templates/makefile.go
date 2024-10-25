package templates

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// Init function called by Go runtime
func init() {
	RegisterSingleFileTemplate("makefile", func() Template {
		return MakefileTemplate{ProjectName: "", Author: ""}
	})
}

// Hold data specific to a Makefile template
type MakefileTemplate struct {
	Author      string
	ProjectName string
}

// Create a Makefile at the specified output path
func (m MakefileTemplate) Generate(outputPath string) error {
	// Read user input
	var author, projectName string

	fmt.Print("Enter author: ")
	_, err := fmt.Scanln(&author)
	if err != nil {
		return err
	}

	fmt.Print("Enter project name: ")
	_, err = fmt.Scanln(&projectName)
	if err != nil {
		return err
	}

	m.Author = author
	m.ProjectName = projectName

	// If outputPath is a directory, join "Makefile" as the default filename
	fileInfo, err := os.Stat(outputPath)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		outputPath = filepath.Join(outputPath, "Makefile")
	}

	// Create or overwrite the output file
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Parse the template
	tmpl, err := template.ParseFiles("templates/Makefile.tmpl")
	if err != nil {
		return err
	}

	// Apply the template with data
	return tmpl.Execute(f, m)
}
