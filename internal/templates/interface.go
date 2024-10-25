package templates

type Template interface {
	Generate(outputPath string) error
}
