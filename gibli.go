package main

import (
	"database/sql"
	"errors"

	"github.com/MihaiBlebea/go-gibli/builder"
	"github.com/MihaiBlebea/go-gibli/migrator"
	"github.com/MihaiBlebea/go-gibli/reader"
	"github.com/MihaiBlebea/go-gibli/reconciliator"
)

// Config contains the details to run the code generator
type Config struct {
	DefinitionsPath string
	ModelsPath      string
	Client          Connect
}

// Connect returns a mysql connection object
type Connect func() *sql.DB

// GenerateModels reads the definition files and builds go models from each
// Impure function
func GenerateModels(c Config) (err error) {
	err = validateConfig(c)
	if err != nil {
		return err
	}

	// Read the yaml definition files for models
	models, err := reader.ReadModelDefinitions(c.DefinitionsPath)
	if err != nil {
		return err
	}

	for _, model := range models {

		// Build model go file
		err = builder.BuildModel(model, c.ModelsPath)
		if err != nil {
			return err
		}

		// Check if the table exist for this model
		err := migrator.CheckTableExists(c.Client(), model.Table)
		if err != nil {
			err = migrator.CreateTable(c.Client(), model)
			if err != nil {
				return err
			}
			continue
		}

		migration, err := reconciliator.Reconciliate(c.Client(), &model)
		if err != nil {
			return err
		}

		if needMigration(migration) {
			migrator.Migrate(c.Client(), migration)
		}
	}
	return nil
}

// GenerateDefinitionFile generates a YAML definition file
func GenerateDefinitionFile(c Config, name, kind string) (err error) {
	err = validateConfig(c)
	if err != nil {
		return err
	}

	err = builder.BuildDefinition(name, kind, "1", c.DefinitionsPath)
	if err != nil {
		return err
	}
	return nil
}

func needMigration(migration builder.Migration) bool {
	if len(migration.Add) > 0 || len(migration.Remove) > 0 || len(migration.Modify) > 0 {
		return true
	}
	return false
}

func validateConfig(c Config) error {
	if c.ModelsPath == "" {
		return errors.New("Invalid config ModelsPath")
	}
	if c.DefinitionsPath == "" {
		return errors.New("Invalid config DefinitionsPath")
	}
	return nil
}
