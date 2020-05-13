package migrator

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/MihaiBlebea/go-gibli/builder"
)

// Migrate migrates a table up
func Migrate(client *sql.DB, migration builder.Migration) error {
	defer client.Close()

	baseSQL := fmt.Sprintf("ALTER TABLE %s", migration.Table)

	var partSQL []string
	if len(migration.Add) > 0 {
		partSQL = append(partSQL, addRows(migration.Add))
	}

	if len(migration.Remove) > 0 {
		partSQL = append(partSQL, dropRows(migration.Remove))
	}

	if len(migration.Modify) > 0 {
		partSQL = append(partSQL, modifyRows(migration.Modify))
	}

	baseSQL += fmt.Sprintf(" %s", strings.Join(partSQL, ", "))

	fmt.Println(baseSQL)

	stmt, err := client.Prepare(baseSQL)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
