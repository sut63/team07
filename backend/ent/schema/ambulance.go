package schema

import (
	"errors"
	"regexp"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Ambulance holds the schema definition for the Ambulance entity.
type Ambulance struct {
	ent.Schema
}

// Fields of the Ambulance.
func (Ambulance) Fields() []ent.Field {
	return []ent.Field{
		field.String("carregistration").
			Unique().NotEmpty().Validate(func(s string) error {
				match, _ := regexp.MatchString("^[0-9]{1}[ก-ฮ]{2}[0-9]{4}$|^[ก-ฮ]{2}[0-9]{4}$", s)
				if !match {
					return errors.New("รูปแบบเลขทะเบียนรถไม่ถูกต้อง")
				}
				return nil
			}),
		field.Int("enginepower").Range(80,100),
		field.Int("displacement").Range(2900,4000),
		field.Time("registerat"),
	}
}

// Edges of the Ambulance.
func (Ambulance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hasbrand", Carbrand.Type).
			Ref("brandof").
			Unique(),
		edge.From("hasinsurance", Insurance.Type).
			Ref("insuranceof").
			Unique(),
		edge.From("hasstatus", InspectionResult.Type).
			Ref("statusof").
			Unique(),
		edge.From("hasuser", User.Type).
			Ref("userof").
			Unique(),

		//To CarInspection
		edge.To("carinspections", CarInspection.Type).
			StorageKey(edge.Column("ambulance_id")),
		//To CarCheckInOut
		edge.To("carcheckinout", CarCheckInOut.Type).
			StorageKey(edge.Column("ambulance")),
		edge.To("ambulance",Transport.Type).
			StorageKey(edge.Column("ambulance")),
	}
}
