package schema

import (
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Send holds the schema definition for the Send entity.
type Send struct {
	ent.Schema
}

// Fields of the Send.
func (Send) Fields() []ent.Field {
	return []ent.Field{
		field.String("sendname").
		Unique(),
	}
}

// Edges of the Send.
func (Send) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sendid",Transport.Type).
		StorageKey(edge.Column("sendid")),
	}
}
