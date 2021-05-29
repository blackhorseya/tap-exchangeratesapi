package singer

// Property declare property struct
type Property struct {
	Type   interface{} `json:"type"`
	Format string      `json:"format"`
}

// Schema declare schema struct
type Schema struct {
	Type       interface{}          `json:"type"`
	Properties map[string]*Property `json:"properties"`
}

// NewSchema serve caller to create default Schema
func NewSchema() *Schema {
	return &Schema{
		Type: "object",
		Properties: map[string]*Property{
			"date": {Type: "string", Format: "date-time"},
		},
	}
}
