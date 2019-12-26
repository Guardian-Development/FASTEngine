package template

const templateTag = "template"

// Template represents an ordered List of operations needed to Serialise/Deserialise a FAST message
type Template struct {
	TemplateUnits []Unit
}
