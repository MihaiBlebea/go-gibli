package builder

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/MihaiBlebea/go-gibli/bundle"
	"github.com/MihaiBlebea/go-gibli/transformer"
)

// BuildModel receives the data, file to write to and template path and writes the model to a go file
func BuildModel(model Model, path string) error {

	functionMap := template.FuncMap{
		"toUpperCase":    strings.Title,
		"toLowerCase":    strings.ToLower,
		"toModelName":    toModelName,
		"toVariableName": toVariableName,
		"toBasicType":    transformer.ToBasicFieldKind,
	}

	err := checkAndCreate(path)
	if err != nil {
		return err
	}

	fileName := extractFileNameFromModel(model.Name)

	templateName := "model.tmpl"
	content := bundle.Get(fmt.Sprintf("templates/%s", templateName))

	tmpl, err := template.New(templateName).Funcs(functionMap).Parse(string(content))
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.go", path, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	model.Module = extractModuleNameFromPath(path)

	err = tmpl.Execute(file, model)
	if err != nil {
		return err
	}
	return nil
}

func toModelName(name string) string {
	if strings.Contains(name, "_") {
		parts := strings.Split(name, "_")
		for index, part := range parts {
			parts[index] = strings.Title(part)
		}
		return strings.Join(parts, "")
	}
	return strings.Title(name)
}

func toVariableName(name string) string {
	if strings.Contains(name, "_") {
		parts := strings.Split(name, "_")
		for index, part := range parts {
			if index == 0 {
				parts[index] = strings.ToLower(part)
			} else {
				parts[index] = strings.Title(part)
			}
		}
		return strings.Join(parts, "")
	}
	return strings.ToLower(name)
}

// pure function
func extractModuleNameFromPath(path string) (moduleName string) {
	last := path[len(path)-1]
	if last == '/' {
		path = path[:last]
	}
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

// pure function
func extractFileNameFromModel(name string) string {
	fileName := strings.ReplaceAll(name, " ", "_")
	fileName = strings.ReplaceAll(fileName, "-", "_")
	return strings.ToLower(fileName)
}
