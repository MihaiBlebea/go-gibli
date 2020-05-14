package builder

import (
	"time"

	"github.com/MihaiBlebea/go-gibli/transformer"
)

// Model is a model
type Model struct {
	Version   string    `yaml:"version"`
	Timestamp time.Time `yaml:"timestamp"`
	Module    string    `yaml:"module"`
	Name      string    `yaml:"name"`
	Kind      string    `yaml:"kind"`
	Table     string    `yaml:"table"`
	Fields    []Field   `yaml:"fields"`
}

// Field is a model
type Field struct {
	Name       string      `yaml:"name"`
	Kind       string      `yaml:"kind"`
	Length     int         `yaml:"length"`
	Searchable bool        `yaml:"searchable"`
	Default    interface{} `yaml:"default"`
	Unique     bool        `yaml:"unique"`
	NotNull    bool        `yaml:"notnull"`
}

// Migration data
type Migration struct {
	Timestamp int64   `yaml:"timestamp"`
	Table     string  `yaml:"table"`
	Add       []Field `yaml:"add"`
	Remove    []Field `yaml:"remove"`
	Modify    []Field `yaml:"modify"`
}

// UnmarshalYAML implements the interface to unmarshal a yaml file into Field struct
func (f *Field) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var field struct {
		Name       string      `yaml:"name"`
		Kind       string      `yaml:"kind"`
		Length     int         `yaml:"length"`
		Searchable bool        `yaml:"searchable"`
		Default    interface{} `yaml:"default"`
		Unique     bool        `yaml:"unique"`
		NotNull    bool        `yaml:"notnull"`
	}
	if err := unmarshal(&field); err != nil {
		return err
	}
	*f = field
	f.Kind = transformer.ToMysqlRowKind(f.Kind)
	return nil
}

// UnmarshalYAML implements the interface to unmarshal a yaml file into Model struct
func (m *Model) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var model struct {
		Version   string    `yaml:"version"`
		Timestamp time.Time `yaml:"timestamp"`
		Module    string    `yaml:"module"`
		Name      string    `yaml:"name"`
		Kind      string    `yaml:"kind"`
		Table     string    `yaml:"table"`
		Fields    []Field   `yaml:"fields"`
	}
	if err := unmarshal(&model); err != nil {
		return err
	}
	*m = model
	m.Table = transformer.ToTableName(m.Name)
	return nil
}
