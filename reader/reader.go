package reader

import (
	"github.com/MihaiBlebea/go-gibli/builder"
)

// ReadModelDefinitions goes into the definition folder, scans all the files
//
// Returns rich structs with the information obtained
//
// **Impure function**
func ReadModelDefinitions(folderPath string) (model []builder.Model, err error) {
	var models []builder.Model

	definitionFiles := scanDefinitionFolder(folderPath)

	// Generate models
	for _, definition := range definitionFiles {
		var model builder.Model

		definitionContent, err := readDefinitionFile(definition)
		if err != nil {
			return models, err
		}

		// Parse the yaml definition files
		m, err := extractModelYaml(definitionContent, &model)
		if err != nil {
			return models, err
		}

		// Validate input from yaml file
		_, err = validateModel(m)
		if err != nil {
			return models, err
		}

		models = append(models, *m)
	}
	return models, nil
}

// ReadConfig reads a YAML config file and returns the values as Config struct
func ReadConfig(path string) (config Config, err error) {
	b, err := readConfigFile(path)
	if err != nil {
		return config, err
	}

	config, err = extractConfigYaml(b, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
