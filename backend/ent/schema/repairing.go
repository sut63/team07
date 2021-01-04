package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
   )
// Repairing holds the schema definition for the Repairing entity.
type Repairing struct {
	ent.Schema
}

// Fields of the Repairing.
func (Repairing) Fields() []ent.Field {
	return []ent.Field{
		field.String("repairpart").NotEmpty().Unique(),
	}
}

// Edges of the Repairing.
func (Repairing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("repairs", CarRepairrecord.Type).
			StorageKey(edge.Column("repairing_id")),
	}
}
