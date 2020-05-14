package migrator

import (
	"database/sql"
	"fmt"

	// Mysql driver for mysql
	_ "github.com/go-sql-driver/mysql"

	"github.com/MihaiBlebea/go-gibli/builder"
)

// CreateTable creates the table
func CreateTable(client *sql.DB, model builder.Model) error {
	defer client.Close()

	rows := createRows(model.Fields)

	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INT NOT NULL AUTO_INCREMENT,
        %s,
        created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted TINYINT(1) DEFAULT 0,
		PRIMARY KEY (id));`, model.Table, rows)

	stmt, err := client.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	return err
}

// RemoveTable removes the table
func RemoveTable(client *sql.DB, tableName string) error {
	defer client.Close()

	sql := fmt.Sprintf(`DROP TABLE IF EXISTS %s;`, tableName)
	stmt, err := client.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	return err
}

// CheckTableExists returns error if table does not exist
func CheckTableExists(client *sql.DB, table string) error {
	defer client.Close()
	sql := fmt.Sprintf("DESCRIBE %s", table)
	_, err := client.Query(sql)
	return err
}
