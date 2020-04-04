package properties

// Properties contains information about a TemplateUnit within a FAST Template
type Properties struct {
	ID       uint64
	Name     string
	Required bool
}

// New properties for a field with the given parameters
func New(id uint64, name string, required bool) Properties {
	props := Properties{
		ID:       id,
		Name:     name,
		Required: required,
	}

	return props
}
