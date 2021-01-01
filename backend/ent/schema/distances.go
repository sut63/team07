package schema
import (
 "github.com/facebookincubator/ent"
 "github.com/facebookincubator/ent/schema/field"
 "github.com/facebookincubator/ent/schema/edge"
)
// Distances holds the schema definition for the Distances entity.
type Distances struct {
 ent.Schema
}
// Fields of the Distances.
func (Distances) Fields() []ent.Field {
 return []ent.Field{
 field.String("Distances").NotEmpty(),
 }
}
// Edges of the Distances.
func (Distances) Edges() []ent.Edge {
 return []ent.Edge{
    edge.To("Disid",Carservice.Type),
 }
}