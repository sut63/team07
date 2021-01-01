// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team07/app/ent/purpose"
)

// Purpose is the model entity for the Purpose schema.
type Purpose struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Purpose) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Purpose fields.
func (pu *Purpose) assignValues(values ...interface{}) error {
	if m, n := len(values), len(purpose.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pu.ID = int(value.Int64)
	values = values[1:]
	return nil
}

// Update returns a builder for updating this Purpose.
// Note that, you need to call Purpose.Unwrap() before calling this method, if this Purpose
// was returned from a transaction, and the transaction was committed or rolled back.
func (pu *Purpose) Update() *PurposeUpdateOne {
	return (&PurposeClient{config: pu.config}).UpdateOne(pu)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pu *Purpose) Unwrap() *Purpose {
	tx, ok := pu.config.driver.(*txDriver)
	if !ok {
		panic("ent: Purpose is not a transactional entity")
	}
	pu.config.driver = tx.drv
	return pu
}

// String implements the fmt.Stringer.
func (pu *Purpose) String() string {
	var builder strings.Builder
	builder.WriteString("Purpose(")
	builder.WriteString(fmt.Sprintf("id=%v", pu.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Purposes is a parsable slice of Purpose.
type Purposes []*Purpose

func (pu Purposes) config(cfg config) {
	for _i := range pu {
		pu[_i].config = cfg
	}
}
