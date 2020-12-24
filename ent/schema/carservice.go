package schema

import "github.com/facebookincubator/ent"

// Carservice holds the schema definition for the Carservice entity.
type Carservice struct {
	ent.Schema
}

// Fields of the Carservice.
func (Carservice) Fields() []ent.Field {
	return nil
}

// Edges of the Carservice.
func (Carservice) Edges() []ent.Edge {
	return nil
}
