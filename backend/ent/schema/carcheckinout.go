package schema

import (
	"time"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// CarCheckInOut holds the schema definition for the CarCheckInOut entity.
type CarCheckInOut struct {
	ent.Schema
}

// Fields of the CarCheckInOut.
func (CarCheckInOut) Fields() []ent.Field {
	return []ent.Field{
		field.String("note"),
		field.Time("checkIn").Default(time.Now),
		field.Time("checkOut"),
	}
}

// Edges of the CarCheckInOut.
func (CarCheckInOut) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ambulance",Ambulance.Type).
			Ref("carcheckinout").
			Unique(),

		edge.From("name",User.Type).
			Ref("carcheckinout").
			Unique(),

		edge.From("purpose",Purpose.Type).
			Ref("carcheckinout").
			Unique(),
		}
}
