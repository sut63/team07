package schema

import "github.com/facebookincubator/ent"

// CarRepairrecord holds the schema definition for the CarRepairrecord entity.
type CarRepairrecord struct {
	ent.Schema
}

// Fields of the CarRepairrecord.
func (CarRepairrecord) Fields() []ent.Field {
	return nil
}

// Edges of the CarRepairrecord.
func (CarRepairrecord) Edges() []ent.Edge {
	return nil
}
