package builder

import (
	"fmt"
	"os"
	"text/template"

	"github.com/MihaiBlebea/go-gibli/bundle"
)

// BuildDefinition creates a new default yaml definition file
func BuildDefinition(name, kind, version, folderPath string) (err error) {
	type Data struct {
		Version string
		Name    string
		Kind    string
	}

	templateName := "definition.tmpl"
	content := bundle.Get(fmt.Sprintf("templates/%s", templateName))

	tmpl, err := template.New(fmt.Sprintf("%s.yaml", name)).Parse(string(content))
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("./%s/%s-definition.yaml", folderPath, name))
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, Data{version, name, kind})
	if err != nil {
		return err
	}
	return nil
}
