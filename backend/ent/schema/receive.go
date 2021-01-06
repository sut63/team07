package schema

import (
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Receive holds the schema definition for the Receive entity.
type Receive struct {
	ent.Schema
}

// Fields of the Receive.
func (Receive) Fields() []ent.Field {
	return []ent.Field{
		field.String("sendname"),
	}
}

// Edges of the Receive.
func (Receive) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("receiveid",Transport.Type).
		StorageKey(edge.Column("receiveid")),
	}
}
