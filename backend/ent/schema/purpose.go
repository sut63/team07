package schema

import "github.com/facebookincubator/ent"

// Purpose holds the schema definition for the Purpose entity.
type Purpose struct {
	ent.Schema
}

// Fields of the Purpose.
func (Purpose) Fields() []ent.Field {
	return nil
}

// Edges of the Purpose.
func (Purpose) Edges() []ent.Edge {
	return nil
}
