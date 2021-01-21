package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Hospital holds the schema definition for the Hospital entity.
type Hospital struct {
	ent.Schema
}

// Fields of the Hospital.
func (Hospital) Fields() []ent.Field {
	return []ent.Field{
		field.String("hospital").Unique(),
	}
}

// Edges of the Hospital.
func (Hospital) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("receive",Transport.Type).
		StorageKey(edge.Column("receive")),
		edge.To("send",Transport.Type).
		StorageKey(edge.Column("send")),
	}
}
