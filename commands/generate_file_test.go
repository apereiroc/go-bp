package commands

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/apereiroc/go-bp/internal/templates"
	"github.com/stretchr/testify/assert"
)

// Private mock implementation of the Template interface for testing
type mockFileTemplate struct{}

func (m *mockFileTemplate) Generate(outputPath string) error {
	// For testing purposes, create a simple output file
	outputPath = filepath.Join(outputPath, "Testfile.test")

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("Mock content")
	return err
}

// Register a mock single-file template for testing
// This mockfile template is only defined when calling go test
func TestRegisterNewFileTemplate_Sucess(t *testing.T) {
	templates.RegisterSingleFileTemplate("mockfile", func() templates.Template {
		return &mockFileTemplate{}
	})

	assert.NoError(t, nil)
}

// Test for the generate-single command's success path
func TestGenerateSingleFileCommand_Success(t *testing.T) {
	// Create temporary directory
	dname, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dname)

	// Set up command and arguments
	cmd := NewFileGeneratorCmd()
	cmd.SetArgs([]string{"mockfile", dname})

	// Run the command
	err = cmd.Execute()

	// Assertions
	assert.NoError(t, err)
}

// Test for the generate-single command with unsupported template type
func TestGenerateSingleFileCommand_NoSuchTemplate(t *testing.T) {
	// Set up command and arguments
	cmd := NewFileGeneratorCmd()
	cmd.SetArgs([]string{"mockfil", "test"})

	// Capture the output
	// var output bytes.Buffer
	// cmd.SetOut(&output)
	// --
	// Run the command
	err := cmd.Execute()

	// Assertions
	assert.Error(t, err)
}

func TestGenerateSingleFileCommand_NoSuchDir(t *testing.T) {
	// Generate a random directory name
	var randomDir string
	for {
		// Generate a random directory name
		randomDir = filepath.Join(os.TempDir(), fmt.Sprintf("nonexistent_dir_%d", rand.Int()))

		// Check if the directory exists
		if _, err := os.Stat(randomDir); os.IsNotExist(err) {
			break
		}
	}

	// Set up command and arguments
	cmd := NewFileGeneratorCmd()
	cmd.SetArgs([]string{"mockfile", randomDir})

	// Capture the output
	// var output bytes.Buffer
	// cmd.SetOut(&output)
	// --
	// Run the command
	err := cmd.Execute()

	// Assertions
	assert.Error(t, err)
}
