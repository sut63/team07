package schema
import (
 "github.com/facebookincubator/ent"
 "github.com/facebookincubator/ent/schema/field"
 "github.com/facebookincubator/ent/schema/edge"
)
// Distance holds the schema definition for the Distance entity.
type Distance struct {
 ent.Schema
}
// Fields of the Distances.
func (Distance) Fields() []ent.Field {
 return []ent.Field{
 field.String("Distance").NotEmpty(),
 }
}
// Edges of the Distances.
func (Distance) Edges() []ent.Edge {
 return []ent.Edge{
    edge.To("disid",Carservice.Type),
 }
}