package transformer

import (
	"fmt"
	"strings"
)

// ToModelName accepts a name as input and returns a kebab case name for variables
//
// pure function
func ToModelName(name string) string {
	return ToFieldName(name)
}

// ToTableName accepts a name as input
//
// Ex: "User", "Car"
//
// Returns a database table name as output
//
// Ex: "users", "cars"
//
// pure function
func ToTableName(name string) string {
	if string(name[len(name)-1]) != "s" {
		name += "s"
	}
	return strings.ToLower(name)
}

// ToFieldName accepts a name as input and returns a kebab case name for variables or database row name
//
// pure function
func ToFieldName(name string) string {
	var parts []string

	parts = strings.Split(name, "_")
	if len(parts) > 1 {
		name = ""
		for _, part := range parts {
			name += strings.Title(part)
		}
	}
	return strings.Title(name)
}

// ToBasicFieldKind accepts a mysql type string as input
//
// Ex: "VARCHAR", "INT", "BIGINT"
//
// Returns a basic type as output
//
// Ex: "string", "bool", "int"
func ToBasicFieldKind(kind string) (typeKind string, err error) {
	stringKind := []string{"CHAR", "VARCHAR", "TEXT", "TINYTEXT", "MEDIUMTEXT", "LONGTEXT", "ENUM"}
	intKind := []string{"INT", "SMALLINT", "MEDIUMINT", "BIGINT", "FLOAT", "DOUBLE", "DECIMAL"}
	dateKind := []string{"DATE", "DATETIME", "TIMESTAMP", "TIME", "YEAR"}
	boolKind := "TINYINT"

	if contains(stringKind, kind) {
		return "string", nil
	}

	if contains(intKind, kind) {
		return "int", nil
	}

	if contains(dateKind, kind) {
		return "date", nil
	}

	if boolKind == kind {
		return "bool", nil
	}
	return "", fmt.Errorf("Could not find type for %s", kind)
}

// ToMysqlRowKind accepts a lower case, camel case or any other style of string format
//
// Ex: "varchar", "tiny-int"
//
// Returns an uppercase string with no "-" or "_"
//
// Ex: "VARCHAR", "TINYINT"
//
// pure function
func ToMysqlRowKind(kind string) string {
	if strings.Contains(kind, " ") {
		kind = strings.ReplaceAll(kind, " ", "")
	}
	if strings.Contains(kind, "-") {
		kind = strings.ReplaceAll(kind, "-", "")
	}
	if strings.Contains(kind, "_") {
		kind = strings.ReplaceAll(kind, "_", "")
	}
	return strings.ToUpper(kind)
}

// func toMysqlFieldName(name string) string {
// 	name = strings.ReplaceAll(name, "-", "_")
// 	return strings.ToLower(name)
// }

// pure function
func contains(in []string, value string) bool {
	for _, item := range in {
		if item == value {
			return true
		}
	}
	return false
}
