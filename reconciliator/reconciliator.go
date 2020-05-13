package reconciliator

import (
	"database/sql"
	"time"

	"github.com/MihaiBlebea/go-gibli/builder"
)

// Reconciliate will put side by side, the new model from the yaml file and the old information from the db for a specific table
//
// If we can find the table in the db it means that this is a migration and not a new table
//
// **NOTE** This function requires a DB connection
// ( *sql.DB )
func Reconciliate(client *sql.DB, model *builder.Model) (migration builder.Migration, err error) {
	tableFields, err := describeTable(client, model.Table)
	if err != nil {
		return migration, err
	}

	migration = compareFields(model.Fields, tableFields)
	migration.Timestamp = time.Now().Unix()
	migration.Table = model.Table

	return migration, nil
}
