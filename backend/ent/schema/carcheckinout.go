package schema

import "github.com/facebookincubator/ent"

// CarCheckInOut holds the schema definition for the CarCheckInOut entity.
type CarCheckInOut struct {
	ent.Schema
}

// Fields of the CarCheckInOut.
func (CarCheckInOut) Fields() []ent.Field {
	return nil
}

// Edges of the CarCheckInOut.
func (CarCheckInOut) Edges() []ent.Edge {
	return nil
}
