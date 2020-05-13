package reader

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MihaiBlebea/go-gibli/builder"
	"github.com/MihaiBlebea/go-gibli/transformer"
)

// TODO
// - Validate if no two fields are the same in the definition file - same name
// - Validate if default value matches the field kind

// pure function
func validateModel(model *builder.Model) (bool, error) {
	var valid bool
	var err error

	valid, err = validateKind(model.Kind)
	if err != nil {
		return valid, err
	}

	for _, field := range model.Fields {
		valid, err = validateFieldName(field.Name)
		if err != nil {
			return valid, err
		}

		valid, err = validateFieldKind(field.Kind)
		if err != nil {
			return valid, err
		}

		valid, err = validateFieldLength(field.Kind, field.Name, field.Length)
		if err != nil {
			return valid, err
		}

		valid, err = validateFieldDefault(field.Kind, field.Default)
		if err != nil {
			return valid, err
		}
	}

	return true, nil
}

// pure function
func validateKind(kind string) (bool, error) {
	valid := []string{"model"}
	if contains(valid, kind) == false {
		return false, fmt.Errorf("Invalid kind %s", kind)
	}
	return true, nil
}

// pure function
func validateFieldName(name string) (bool, error) {
	if strings.Contains(name, " ") {
		return false, fmt.Errorf("Invalid field name %s", name)
	}
	if strings.Contains(name, "-") {
		return false, fmt.Errorf("Invalid field name %s", name)
	}
	return true, nil
}

// pure function
func validateFieldLength(kind, name string, length int) (bool, error) {
	max, found := fieldMaxLength(kind)
	if found == false {
		return true, nil
	}
	if length <= max {
		return true, nil
	}
	return false, fmt.Errorf("Field %s is too long", name)
}

// pure function
func validateFieldDefault(kind string, value interface{}) (bool, error) {
	var ok bool

	typeKind, err := transformer.ToBasicFieldKind(kind)
	if err != nil {
		return false, err
	}

	switch typeKind {
	case "string", "date":
		_, ok = value.(string)
		return ok, nil
	case "int":
		_, ok = value.(int)
		return ok, nil
	case "bool":
		_, ok = value.(bool)
		return ok, nil
	}
	return false, errors.New("Could not find field kind")
}

// pure function
func fieldMaxLength(kind string) (int, bool) {
	switch kind {
	case "INT":
		return 11, true
	case "TINYINT":
		return 4, true
	case "SMALLINT":
		return 5, true
	case "MEDIUMINT":
		return 9, true
	case "BIGINT":
		return 20, true
	case "FLOAT":
		return 24, true
	case "DOUBLE":
		return 53, true
	case "CHAR":
		return 255, true
	case "VARCHAR":
		return 255, true
	}
	return 0, false
}

// pure function
func contains(in []string, value string) bool {
	for _, item := range in {
		if item == value {
			return true
		}
	}
	return false
}

func validateFieldKind(kind string) (bool, error) {
	valid := []string{
		"INT", "TINYINT", "SMALLINT", "MEDIUMINT", "BIGINT", "FLOAT", "DOUBLE", "DECIMAL", "DATE", "DATETIME", "TIMESTAMP",
		"TIME", "YEAR", "CHAR", "VARCHAR", "TEXT", "TINYTEXT", "MEDIUMTEXT", "LONGTEXT", "ENUM",
	}
	if contains(valid, kind) == false {
		return false, fmt.Errorf("Invalid field kind for %s", kind)
	}
	return true, nil
}

// https://www.tutorialspoint.com/mysql/mysql-data-types.htm
