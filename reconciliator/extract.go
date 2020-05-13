package reconciliator

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/MihaiBlebea/go-gibli/builder"
	"github.com/MihaiBlebea/go-gibli/transformer"
)

type tableCol struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default *string
	Extra   string
}

func describeTable(client *sql.DB, tableName string) (fieldCollection []builder.Field, err error) {
	defer client.Close()

	rows, err := client.Query(fmt.Sprintf("DESCRIBE %s", tableName))
	if err != nil {
		return fieldCollection, err
	}

	for rows.Next() {
		var col tableCol

		err = rows.Scan(
			&col.Field,
			&col.Type,
			&col.Null,
			&col.Key,
			&col.Default,
			&col.Extra,
		)
		if err != nil {
			return fieldCollection, err
		}

		// Do not add default fields
		if col.Field != "id" && col.Field != "created" && col.Field != "updated" && col.Field != "deleted" {
			field, err := fromTableColToField(col)
			if err != nil {
				return fieldCollection, err
			}
			fieldCollection = append(fieldCollection, field)
		}
	}
	return fieldCollection, nil
}

func fromTableColToField(row tableCol) (field builder.Field, err error) {
	field.Name = row.Field

	kind, length, err := extractKindAndLengthFromRow(row.Type)
	if err != nil {
		return field, err
	}
	field.Kind = kind
	field.Length = length
	field.NotNull = notNullToBool(row.Null)
	field.Unique = isUnique(row.Key)

	if hasDefault(row.Default) {
		def := extractStringFromPointer(row.Default)
		field.Default, err = extractDefault(def, kind)
		if err != nil {
			return field, err
		}
	}

	return field, nil
}

// Transforms field kind from mysql format (VARCHAR, INT) to go type format (string, int)
//
// Receives field kind in mysql format
//
// Returns field kind in go type format and length of the field
//
// **NOTE** Pure function
func extractKindAndLengthFromRow(kind string) (k string, lenght int, err error) {
	if strings.Contains(kind, "(") {
		parts := strings.FieldsFunc(kind, Split)
		if len(parts) != 2 {
			return "", 0, errors.New("Invalid number of parts from string split")
		}
		length, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", 0, err
		}
		return strings.ToUpper(parts[0]), length, nil
	}
	return strings.ToUpper(kind), 0, nil
}

// Split custom split function for strings
// Pure function
func Split(r rune) bool {
	return r == '(' || r == ')'
}

func notNullToBool(value string) bool {
	if value == "NO" {
		return true
	}
	return false
}

func isUnique(key string) bool {
	if key == "UNI" {
		return true
	}
	return false
}

func hasDefault(value *string) bool {
	if value == nil {
		return false
	}
	return true
}

func extractDefault(value, kind string) (result interface{}, err error) {
	basicKind, err := transformer.ToBasicFieldKind(kind)

	if err != nil {
		return result, err
	}
	if basicKind == "string" {
		return value, nil
	}

	if basicKind == "int" {
		integer, err := strconv.Atoi(value)
		if err != nil {
			return result, err
		}
		return integer, nil
	}

	if basicKind == "bool" {
		if value == "1" {
			return true, nil
		}
		return false, nil
	}
	return value, nil
}

func extractStringFromPointer(value *string) string {
	if value == nil {
		return ""
	}
	val := *value
	return val
}
