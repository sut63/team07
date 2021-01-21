package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// CarInspection holds the schema definition for the CarInspection entity.
type CarInspection struct {
	ent.Schema
}

// Fields of the CarInspection.
func (CarInspection) Fields() []ent.Field {
	return []ent.Field{
		field.Float("wheel_center").Positive(),
		field.Float("sound_level").Positive(),
		field.Float("blacksmoke").Positive().Max(100),

		field.Time("datetime"),
		field.String("note"),
	}
}

// Edges of the CarInspection.
func (CarInspection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("carinspections").
			Unique(),
		edge.From("ambulance", Ambulance.Type).
			Ref("carinspections").
			Unique(),
		edge.From("inspectionresult", InspectionResult.Type).
			Ref("carinspections").
			Unique(),
		//To CarRepairrecord
		edge.To("carrepairrecords", CarRepairrecord.Type).
			StorageKey(edge.Column("carinspection_id")),
	}
}
