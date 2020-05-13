package migrator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MihaiBlebea/go-gibli/builder"
)

// pure function
func createRow(field builder.Field) string {
	var base string
	length := strconv.Itoa(field.Length)
	if field.Length > 0 {
		base = fmt.Sprintf("%s %s(%s)", field.Name, field.Kind, length)
	} else {
		base = fmt.Sprintf("%s %s", field.Name, field.Kind)
	}

	// set options
	var options []string
	if field.NotNull == true {
		options = append(options, "NOT NULL")
	}

	if field.Unique == true {
		options = append(options, "UNIQUE")
	}

	if field.Default != "" {
		if value, ok := field.Default.(int); ok {
			options = append(options, fmt.Sprintf("DEFAULT %s", strconv.Itoa(value)))
		}
		if value, ok := field.Default.(string); ok {
			options = append(options, fmt.Sprintf("DEFAULT '%s'", value))
		}
		if value, ok := field.Default.(bool); ok {
			options = append(options, fmt.Sprintf("DEFAULT %s", strconv.FormatBool(value)))
		}
	}
	return fmt.Sprintf("%s %s", base, strings.Join(options, " "))
}

// pure function
func createRows(fields []builder.Field) string {
	var rows []string
	for _, field := range fields {
		row := createRow(field)
		rows = append(rows, row)
	}

	return strings.Join(rows, ", ")
}

// pure function
func addRow(field builder.Field) string {
	return fmt.Sprintf("ADD %s", createRow(field))
}

// pure function
func addRows(fields []builder.Field) string {
	var rows []string
	for _, field := range fields {
		row := addRow(field)
		rows = append(rows, row)
	}
	return strings.Join(rows, ", ")
}

// pure function
func dropRow(field builder.Field) string {
	return fmt.Sprintf("DROP COLUMN %s", field.Name)
}

// pure function
func dropRows(fields []builder.Field) string {
	var rows []string
	for _, field := range fields {
		row := dropRow(field)
		rows = append(rows, row)
	}
	return strings.Join(rows, ", ")
}

func addUniqueIndex(columnName string) string {
	return fmt.Sprintf("ADD UNIQUE (%s)", columnName)
}

func dropUniqueIndex(columnName string) string {
	return fmt.Sprintf("DROP INDEX %s", columnName)
}

func modifyRow(field builder.Field) string {
	return fmt.Sprintf("MODIFY COLUMN %s", createRow(field))
	// if field.Default == false {
	// 	modify = append(modify, dropUniqueIndex(field.Name))
	// } else {
	// 	modify = append(modify, addUniqueIndex(field.Name))
	// }
	// return strings.Join(modify, ", ")
}

func modifyRows(fields []builder.Field) string {
	var rows []string
	for _, field := range fields {
		row := modifyRow(field)
		rows = append(rows, row)
	}
	return strings.Join(rows, ", ")
}
