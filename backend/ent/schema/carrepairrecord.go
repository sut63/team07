package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
	"errors"
	"regexp"
)

// CarRepairrecord holds the schema definition for the CarRepairrecord entity.
type CarRepairrecord struct {
	ent.Schema
}

// Fields of the CarRepairrecord.
func (CarRepairrecord) Fields() []ent.Field {
	return []ent.Field{
		field.Time("datetime"),
		field.String("partrepair").Validate(func(d string) error {
			match, _ := regexp.MatchString("^[ก-๙a-zA-Z-\\s]+$", d)
			if !match {
				return errors.New("รูปแบบส่วนที่ซ่อมไม่ถูกต้อง")
			}
			return nil
		}),
		field.Int("price").Positive(),
		field.String("techniciancomment").Validate(func(s string) error {
				match, _ := regexp.MatchString("^[ก-๙0-9a-zA-Z-\\s]+$", s)
				if !match {
					return errors.New("รูปแบบคอมเมนท์ไม่ถูกต้อง")
				}
				return nil
			}),
	}
}

// Edges of the CarRepairrecord.
func (CarRepairrecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("keeper", Repairing.Type).
			Ref("repairs").Unique(),
		edge.From("user", User.Type).
			Ref("carrepairrecords").Unique(),
		edge.From("carinspection", CarInspection.Type).
			Ref("carrepairrecords").Unique(),
	}
}
