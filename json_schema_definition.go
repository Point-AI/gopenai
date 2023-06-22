package gopenai

const (
	JSONSchemaObject  = "object"
	JSONSchemaNumber  = "number"
	JSONSchemaString  = "string"
	JSONSchemaArray   = "array"
	JSONSchemaNull    = "null"
	JSONSchemaBoolean = "boolean"
)

type JSONSchemaDefinition struct {
	Type        string                          `json:"type,omitempty"`
	Description string                          `json:"description,omitempty"`
	Enum        []string                        `json:"enum,omitempty"`
	Properties  map[string]JSONSchemaDefinition `json:"properties,omitempty"`
	Required    []string                        `json:"required,omitempty"`
	Items       *JSONSchemaDefinition           `json:"items,omitempty"`
}
