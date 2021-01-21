package schema

import (
	//"github.com/facebookincubator/ent/schema/field"
	"regexp"
	"errors"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Transport holds the schema definition for the Transport entity.
type Transport struct {
	ent.Schema
}

// Fields of the Transport.
func (Transport) Fields() []ent.Field {
	return []ent.Field{
		field.String("symptom").NotEmpty().
		Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๙0-9a-zA-Z\\s-]+$", s)
			if !match {
				return errors.New("กรุณณาใส่ข้อมูลให้ถูกต้อง ถ้าไม่มีให้ใส่ - ")
			}
			return nil
		}),
		field.String("drugallergy").NotEmpty().
		Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๙0-9a-zA-Z\\s-]+$", s)
			if !match {
				return errors.New("กรุณณาใส่ข้อมูลให้ถูกต้อง ถ้าไม่มีให้ใส่ - ")
			}
			return nil
		}),
		field.String("note").NotEmpty().
		Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๙0-9a-zA-Z\\s-]+$", s)
			if !match {
				return errors.New("กรุณณาใส่หมายเหตุเพิ่มเติม ถ้าไม่มีให้ใส่ - ")
			}
			return nil
		}),
	}
}

// Edges of the Transport.
func (Transport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("send",Hospital.Type).
			Ref("send").Unique(),
		edge.From("receive",Hospital.Type).
			Ref("receive").Unique(),

		edge.From("user", User.Type).
			Ref("user").Unique(),

		edge.From("ambulance",Ambulance.Type).
			Ref("ambulance").Unique(),
			
	}
}
