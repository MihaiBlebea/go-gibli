package reader

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/MihaiBlebea/go-gibli/builder"
	"gopkg.in/yaml.v2"
)

// impure function
func scanDefinitionFolder(path string) []string {
	var files []string
	filepath.Walk(filepath.Clean(path), func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files
}

// impure function
func readDefinitionFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// pure function
func extractModelYaml(content []byte, model *builder.Model) (*builder.Model, error) {
	err := yaml.Unmarshal(content, &model)
	if err != nil {
		return model, err
	}
	return model, nil
}
