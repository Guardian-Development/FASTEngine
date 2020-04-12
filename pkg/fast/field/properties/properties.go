package properties

import "log"

// Properties contains information about a TemplateUnit within a FAST Template
type Properties struct {
	ID       uint64
	Name     string
	Required bool

	Logger *log.Logger
}

// New properties for a field with the given parameters
func New(id uint64, name string, required bool, logger *log.Logger) Properties {
	props := Properties{
		ID:       id,
		Name:     name,
		Required: required,
		Logger:   logger,
	}

	return props
}
