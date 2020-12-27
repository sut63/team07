// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team07/app/ent/servicetype"
)

// Servicetype is the model entity for the Servicetype schema.
type Servicetype struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Servicetype) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Servicetype fields.
func (s *Servicetype) assignValues(values ...interface{}) error {
	if m, n := len(values), len(servicetype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	return nil
}

// Update returns a builder for updating this Servicetype.
// Note that, you need to call Servicetype.Unwrap() before calling this method, if this Servicetype
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Servicetype) Update() *ServicetypeUpdateOne {
	return (&ServicetypeClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Servicetype) Unwrap() *Servicetype {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Servicetype is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Servicetype) String() string {
	var builder strings.Builder
	builder.WriteString("Servicetype(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Servicetypes is a parsable slice of Servicetype.
type Servicetypes []*Servicetype

func (s Servicetypes) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
