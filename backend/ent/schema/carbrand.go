package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)
// Carbrand holds the schema definition for the Carbrand entity.
type Carbrand struct {
	ent.Schema
}

// Fields of the Carbrand.
func (Carbrand) Fields() []ent.Field {
	return []ent.Field{
		field.String("brand").
		Unique(),
	}
}

// Edges of the Carbrand.
func (Carbrand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("brandof",Ambulance.Type).
		   StorageKey(edge.Column("brand_id")),
	}
}
