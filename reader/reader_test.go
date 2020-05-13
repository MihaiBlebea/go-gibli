package reader

import (
	"strconv"
	"testing"
)

// Tests
func TestValidateKind(t *testing.T) {
	var ios = map[string]bool{
		"model": true,
		"foo":   false,
		"bar":   false,
	}
	for input, output := range ios {
		out, _ := validateKind(input)
		if out != output {
			t.Errorf("Should return %s for input %s", input, strconv.FormatBool(out))
		}
	}
}

func TestValidateFieldName(t *testing.T) {
	var ios = map[string]bool{
		"user":      true,
		"username":  true,
		"user_name": true,
		"user-name": false,
		"User-Name": false,
		"UserName":  true,
	}
	for input, output := range ios {
		out, _ := validateFieldName(input)
		if out != output {
			t.Errorf("Should return %s for input %s", input, strconv.FormatBool(out))
		}
	}
}

func TestValidateFieldKind(t *testing.T) {
	var ios = map[string]bool{
		"string": true,
		"int":    true,
		"bool":   true,
		"slice":  false,
		"object": false,
		"String": false,
	}
	for input, output := range ios {
		out, _ := validateFieldKind(input)
		if out != output {
			t.Errorf("Should return %s for input %s", input, strconv.FormatBool(out))
		}
	}
}

func TestValidateFieldLength(t *testing.T) {
	type io struct {
		kind   string
		name   string
		length int
		output bool
	}
	var ios = []io{
		io{"string", "username", 23, true},
		io{"int", "age", 11, true},
		io{"bool", "is_free", 1, true},
		io{"bool", "is_pink", 2, false},
		io{"int", "money", 23000, false},
		io{"string", "job", 23000, false},
	}
	for _, io := range ios {
		out, _ := validateFieldLength(io.kind, io.name, io.length)
		if out != io.output {
			t.Errorf("Should return %s for %s of length %s", strconv.FormatBool(io.output), io.kind, strconv.Itoa(io.length))
		}
	}
}

// Benchmarks
func BenchmarkValidateKind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validateKind("model")
	}
}

func BenchmarkValidateFieldName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validateFieldName("user_name")
	}
}

func BenchmarkValidateFieldKind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validateFieldKind("string")
	}
}

func BenchmarkValidateFieldLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validateFieldLength("string", "foo_bar", 100)
	}
}
