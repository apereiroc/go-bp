package templates

import (
	"os"
	"path/filepath"
	"text/template"
)

// Hold data specific to a Makefile template
type MakefileTemplate struct {
	Author      string
	ProjectName string
}

// Create a Makefile at the specified output path
func (m MakefileTemplate) Generate(outputPath string) error {
	// If outputPath is a directory, join "Makefile" as the default filename
	fileInfo, err := os.Stat(outputPath)
	if err == nil && fileInfo.IsDir() {
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
