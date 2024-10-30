package templates

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakefileTemplate_Generate(t *testing.T) {
	// Set up a mock MakefileTemplate with test values
	template := MakefileTemplate{}
	template.SetAuthorAndProject("Test Author", "Test Project")

	// Use a temporary directory for output
	outputPath := t.TempDir()

	// Generate the Makefile
	err := template.Generate(outputPath)
	assert.NoError(t, err)

	// fileInfo, err := os.Executable()
	// assert.NoError(t, err)
	// t.Log("fileInfo: ", fileInfo)

	// Verify the file was created
	_, err = os.Stat(outputPath)
	assert.NoError(t, err)

	// Check file contents to confirm expected template rendering
	outputFile := filepath.Join(outputPath, "Makefile")
	data, err := os.ReadFile(outputFile)
	assert.NoError(t, err)
	dataStr := string(data)
	assert.Contains(t, dataStr, "Test Author")
	assert.Contains(t, dataStr, "Test Project")
}
