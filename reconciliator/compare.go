package reconciliator

import (
	"fmt"

	"github.com/MihaiBlebea/go-gibli/builder"
)

// Returns bool if the two fields are the same
//
// pure function
func compareFieldAttributes(field builder.Field, f builder.Field) (same bool) {

	fmt.Println(field)
	fmt.Println(f)

	if field.Name != f.Name {
		return false
	}

	if field.Kind != f.Kind {
		return false
	}

	if field.Length != f.Length {
		return false
	}

	if field.NotNull != f.NotNull {
		return false
	}

	if field.Default != f.Default {
		return false
	}

	// if field.Unique != f.Unique {
	// 	return false
	// }

	return true
}

func compareFields(newFields, oldFields []builder.Field) (migration builder.Migration) {
	for _, field := range newFields {
		var found bool
		for index, f := range oldFields {
			if f.Name == field.Name {
				found = true
				if compareFieldAttributes(field, f) == false {
					migration.Modify = append(migration.Modify, field)
				}
				oldFields = removeFromCollection(oldFields, index)
				break
			}
		}

		if found == false {
			migration.Add = append(migration.Add, field)
		}
	}

	migration.Remove = oldFields

	return migration
}

func removeFromCollection(collection []builder.Field, index int) []builder.Field {
	return append(collection[:index], collection[index+1:]...)
}
