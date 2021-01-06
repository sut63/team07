package schema

import (
	//"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Transport holds the schema definition for the Transport entity.
type Transport struct {
	ent.Schema
}

// Fields of the Transport.
func (Transport) Fields() []ent.Field {
	return nil
}

// Edges of the Transport.
func (Transport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sendid",Send.Type).
			Ref("sendid").Unique(),
		edge.From("receiveid",Receive.Type).
			Ref("receiveid").Unique(),

		edge.From("user", User.Type).
			Ref("user").Unique(),

		edge.From("ambulance",Ambulance.Type).
			Ref("ambulance").Unique(),
			
	}
}
