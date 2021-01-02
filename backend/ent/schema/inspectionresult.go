package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// InspectionResult holds the schema definition for the InspectionResult entity.
type InspectionResult struct {
	ent.Schema
}

// Fields of the InspectionResult.
func (InspectionResult) Fields() []ent.Field {
	return []ent.Field{
		field.String("result_name").NotEmpty().Unique(),
	}
}

// Edges of the InspectionResult.
func (InspectionResult) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("carinspections",CarInspection.Type).
			StorageKey(edge.Column("inspectionresult_id")),
		edge.To("statusof", Ambulance.Type).
			StorageKey(edge.Column("carstatus_id")),
	}
}
