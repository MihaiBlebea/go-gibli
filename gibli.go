package gibli

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/MihaiBlebea/go-gibli/builder"
	"github.com/MihaiBlebea/go-gibli/migrator"
	"github.com/MihaiBlebea/go-gibli/reader"
	"github.com/MihaiBlebea/go-gibli/reconciliator"
)

// Config contains the details to run the code generator
type Config struct {
	ModelDefinitionPath string
	ModelFilesPath      string
	Client              Connect
}

// Connect returns a mysql connection object
type Connect func() *sql.DB

// GenerateModels reads the definition files and builds go models from each
// Impure function
func GenerateModels(c Config) error {
	if c.ModelDefinitionPath == "" {
		return errors.New("Invalid config ModelDefinitionPath")
	}
	if c.ModelFilesPath == "" {
		return errors.New("Invalid config ModelFilesPath")
	}

	// Read the yaml definition files for models
	models, err := reader.ReadModelDefinitions(c.ModelDefinitionPath)
	if err != nil {
		return err
	}

	for _, model := range models {

		// Build model go file
		err = builder.BuildModel(model, c.ModelFilesPath)
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

		fmt.Println(migration)

		if needMigration(migration) {
			migrator.Migrate(c.Client(), migration)
		}
	}
	return nil
}

func needMigration(migration builder.Migration) bool {
	if len(migration.Add) > 0 || len(migration.Remove) > 0 || len(migration.Modify) > 0 {
		return true
	}
	return false
}
