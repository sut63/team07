package schema

import (
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Carservice holds the schema definition for the Carservice entity.
type Carservice struct {
	ent.Schema
}

// Fields of the Carservice.
func (Carservice) Fields() []ent.Field {
	return []ent.Field{
		field.String("customer"),
		field.String("location"),
		field.Time("Datetime"),
	}
}

// Edges of the Carservice.
func (Carservice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userid",User.Type).
			Ref("userid").Unique(),
		edge.From("disid",Distance.Type).
			Ref("disid").Unique(),
		edge.From("urgentid",Urgent.Type).
			Ref("urgentid").Unique(),
	}
}
