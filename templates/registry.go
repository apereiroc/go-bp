package templates

import "fmt"

// Keep track of available templates
var singleFileTemplates = map[string]func() Template{}

// Register a new single-file template.
func RegisterSingleFileTemplate(name string, constructor func() Template) {
	singleFileTemplates[name] = constructor
}

// Get a single-file template by name.
func GetSingleFileTemplate(name string) (Template, error) {
	if constructor, exists := singleFileTemplates[name]; exists {
		return constructor(), nil
	}
	return nil, fmt.Errorf("template type '%s' is not supported", name)
}

// Returns the names of all supported templates
func ListTemplates() (single []string) {
	for k := range singleFileTemplates {
		single = append(single, k)
	}
	return single
}
