package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Ambulance holds the schema definition for the Ambulance entity.
type Ambulance struct {
	ent.Schema
}

// Fields of the Ambulance.
func (Ambulance) Fields() []ent.Field {
	return []ent.Field{
		field.String("carregistration").
			Unique(),
		field.Time("registerat").
			Default(time.Now),
	}
}

// Edges of the Ambulance.
func (Ambulance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hasbrand", Carbrand.Type).
			Ref("brandof").
			Unique(),
		edge.From("hasinsurance", Insurance.Type).
			Ref("insuranceof").
			Unique(),
		edge.From("hasstatus", InspectionResult.Type).
			Ref("statusof").
			Unique(),
		edge.From("hasuser", User.Type).
			Ref("userof").
			Unique(),

		//To CarInspection
		edge.To("carinspections", CarInspection.Type).
			StorageKey(edge.Column("ambulance_id")),
	}
}
