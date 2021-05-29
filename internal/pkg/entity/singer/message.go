package singer

// SchemaMessage declare schema message
type SchemaMessage struct {
	Type   string  `json:"type"`
	Stream string  `json:"stream"`
	Schema *Schema `json:"schema"`
	KeyProperties []string `json:"key_properties"`
	Selected bool `json:"selected"`
}

// NewSchemaMessage serve caller to create a SchemaMessage
func NewSchemaMessage(stream string, schema *Schema, keyProperties []string) *SchemaMessage {
	return &SchemaMessage{
		Type:   "SCHEMA",
		Stream: stream,
		Schema: schema,
		KeyProperties: keyProperties,
	}
}
