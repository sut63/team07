package schema
import (
 "github.com/facebookincubator/ent"
 "github.com/facebookincubator/ent/schema/field"
 "github.com/facebookincubator/ent/schema/edge"
)
// Range holds the schema definition for the Range entity.
type Range struct {
 ent.Schema
}
// Fields of the Range.
func (Range) Fields() []ent.Field {
 return []ent.Field{
 field.String("range").NotEmpty(),
 }
}
// Edges of the Range.
func (Range) Edges() []ent.Edge {
 return []ent.Edge{
    edge.To("rangeid",Carservice.Type),
 }
}