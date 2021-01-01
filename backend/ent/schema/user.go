package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}
//
// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("jobposition", JobPosition.Type).
			Ref("users").
			Unique(),
		edge.To("userof",Ambulance.Type).
			StorageKey(edge.Column("user_id")),
		edge.To("userid",Carservice.Type).
			StorageKey(edge.Column("user_id")),
			
		//To CarInspection	
		edge.To("carinspections",CarInspection.Type).
		    StorageKey(edge.Column("user_id")),
		    
	}
}
