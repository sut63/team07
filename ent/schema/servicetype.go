package schema

import "github.com/facebookincubator/ent"

// Servicetype holds the schema definition for the Servicetype entity.
type Servicetype struct {
	ent.Schema
}

// Fields of the Servicetype.
func (Servicetype) Fields() []ent.Field {
	return nil
}

// Edges of the Servicetype.
func (Servicetype) Edges() []ent.Edge {
	return nil
}
