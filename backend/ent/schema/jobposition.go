package schema

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// JobPosition holds the schema definition for the JobPosition entity.
type JobPosition struct {
	ent.Schema
}

// Fields of the JobPosition.
func (JobPosition) Fields() []ent.Field {
	return []ent.Field{
		field.String("position_name").NotEmpty().Unique(),
	}
}

// Edges of the JobPosition.
func (JobPosition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users",User.Type).StorageKey(edge.Column("jobposition_id")),
	}
}
