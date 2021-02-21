package schema

import (
	"regexp"
	"errors"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Carservice holds the schema definition for the Carservice entity.
type Carservice struct {
	ent.Schema
}

// Fields of the Carservice.
func (Carservice) Fields() []ent.Field {
	return []ent.Field{
		field.String("customer").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๙a-zA-Z-\\s]+$",s)
			if !match {
				return errors.New("รูปแบบไม่ถูกต้อง")
			}
			return nil
		}),
		field.Int("age").Positive(),

		field.String("location").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๙0-9a-zA-Z-\\s]+$",s)
			if !match {
				return errors.New("รูปแบบไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("serviceinfo").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๙0-9a-zA-Z-\\s]+$",s)
			if !match {
				return errors.New("รูปแบบไม่ถูกต้อง")
			}
			return nil
		}),
		field.Time("Datetime"),
	}
}

// Edges of the Carservice.
func (Carservice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("userid",User.Type).
			Ref("userid").Unique(),
		edge.From("disid",Distance.Type).
			Ref("disid").Unique(),
		edge.From("urgentid",Urgent.Type).
			Ref("urgentid").Unique(),
	}
}
