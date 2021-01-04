package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// CarRepairrecord holds the schema definition for the CarRepairrecord entity.
type CarRepairrecord struct {
	ent.Schema
}

// Fields of the CarRepairrecord.
func (CarRepairrecord) Fields() []ent.Field {
	return []ent.Field{
		field.Time("datetime"),
	}
}

// Edges of the CarRepairrecord.
func (CarRepairrecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("keeper", Repairing.Type).
			Ref("repairs").Unique(),
		edge.From("user", User.Type).
			Ref("carrepairrecords").Unique(),
		edge.From("carinspection", CarInspection.Type).
			Ref("carrepairrecords").Unique(),
	}
}
