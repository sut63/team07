package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)
// Insurance holds the schema definition for the Insurance entity.
type Insurance struct {
	ent.Schema
}

// Fields of the Insurance.
func (Insurance) Fields() []ent.Field {
	return []ent.Field{
		field.String("classofinsurance"),
		field.String("company"),
	}
}

// Edges of the Insurance.
func (Insurance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("insuranceof",Ambulance.Type).
		   StorageKey(edge.Column("insurance_id")),
	}
}
