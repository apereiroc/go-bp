package templates

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func init() {
	RegisterSingleFileTemplate("makefile", func() Template {
		return MakefileTemplate{ProjectName: "", Author: ""}
	})
}

type MakefileTemplate struct {
	Author      string
	ProjectName string
}

func (m *MakefileTemplate) SetAuthorAndProject(author, projectName string) {
	m.Author = author
	m.ProjectName = projectName
}

//go:embed static/Makefile.tmpl
var makefileTemplateContent embed.FS

func (m MakefileTemplate) Generate(outputPath string) error {
	// If author and project name are not set, prompt user for them
	if m.Author == "" && m.ProjectName == "" {
		fmt.Print("Enter author: ")
		_, err := fmt.Scanln(&m.Author)
		if err != nil {
			return err
		}

		fmt.Print("Enter project name: ")
		_, err = fmt.Scanln(&m.ProjectName)
		if err != nil {
			return err
		}
	}

	// If outputPath is a directory, join "Makefile" as the default filename
	fileInfo, err := os.Stat(outputPath)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		outputPath = filepath.Join(outputPath, "Makefile")
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl, err := template.ParseFS(makefileTemplateContent, "static/Makefile.tmpl")
	if err != nil {
		return err
	}

	return tmpl.Execute(f, m)
}
