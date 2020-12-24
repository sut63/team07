package schema
import (
 "github.com/facebookincubator/ent"
 "github.com/facebookincubator/ent/schema/field"
 "github.com/facebookincubator/ent/schema/edge"
)
// Urgent holds the schema definition for the Urgent entity.
type Range struct {
 ent.Schema
}
// Fields of the Urgent.
func (Urgent) Fields() []ent.Field {
 return []ent.Field{
 field.String("urgent").NotEmpty(),
 }
}
// Edges of the Urgent.
func (Urgent) Edges() []ent.Edge {
 return []ent.Edge{
    edge.To("urgentid",Carservice.Type),
 }
}