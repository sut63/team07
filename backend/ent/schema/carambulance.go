package schema

import "github.com/facebookincubator/ent"

// CarAmbulance holds the schema definition for the CarAmbulance entity.
type CarAmbulance struct {
	ent.Schema
}

// Fields of the CarAmbulance.
func (CarAmbulance) Fields() []ent.Field {
	return nil
}

// Edges of the CarAmbulance.
func (CarAmbulance) Edges() []ent.Edge {
	return nil
}
