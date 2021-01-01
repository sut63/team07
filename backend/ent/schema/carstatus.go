package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)
// Carstatus holds the schema definition for the Carstatus entity.
type Carstatus struct {
	ent.Schema
}

// Fields of the Carstatus.
func (Carstatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("carstatus").
		Unique(),
	}
}

// Edges of the Carstatus.
func (Carstatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("statusof",Ambulance.Type).
		StorageKey(edge.Column("status_id")),
			
		}
}

