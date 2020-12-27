package schema
import (
 "github.com/facebookincubator/ent"
 "github.com/facebookincubator/ent/schema/field"
 "github.com/facebookincubator/ent/schema/edge"
)
// Servicetype holds the schema definition for the Servicetype entity.
type Servicetype struct {
 ent.Schema
}
// Fields of the Servicetype.
func (Servicetype) Fields() []ent.Field {
 return []ent.Field{
 	field.String("servicetype").NotEmpty(),
 	field.String("servicetype").NotEmpty(),
 }
}
// Edges of the Servicetype.
func (Servicetype) Edges() []ent.Edge {
 return []ent.Edge{
    edge.To("typeid",Carservice.Type),
 }
}