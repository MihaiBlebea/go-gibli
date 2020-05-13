package builder

import (
	"fmt"
	"os"
	"text/template"
)

// BuildDefinition creates a new default yaml definition file
func BuildDefinition(name, kind, version, folderPath string) error {
	templatePath := "./orm/templates/definition.tmpl"

	type Data struct {
		Version string
		Name    string
		Kind    string
	}

	tmpl, err := template.New(fmt.Sprintf("%s.yaml", name)).ParseFiles(templatePath)
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
