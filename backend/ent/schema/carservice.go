package schema

import "github.com/facebookincubator/ent"

// Carservice holds the schema definition for the Carservice entity.
type Carservice struct {
	ent.Schema
}

// Fields of the Carservice.
func (Carservice) Fields() []ent.Field {
	return []ent.Field{
		field.String("customer"),
		field.String("location"),
	}
}

// Edges of the Carservice.
func (Carservice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userid",User.Type).
			Ref("userid").Unique(),
		edge.From("typeid",Servicetype.Type).
			Ref("typeid").Unique(),
		edge.From("rangeid",Range.Type).
			Ref("rangeid").Unique(),
		edge.From("urgentid",Urgent.Type).
			Ref("urgentid").Unique(),
	}
}