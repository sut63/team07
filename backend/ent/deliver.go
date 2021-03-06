// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team07/app/ent/deliver"
)

// Deliver is the model entity for the Deliver schema.
type Deliver struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Deliver) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Deliver fields.
func (d *Deliver) assignValues(values ...interface{}) error {
	if m, n := len(values), len(deliver.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	return nil
}

// Update returns a builder for updating this Deliver.
// Note that, you need to call Deliver.Unwrap() before calling this method, if this Deliver
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Deliver) Update() *DeliverUpdateOne {
	return (&DeliverClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Deliver) Unwrap() *Deliver {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Deliver is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Deliver) String() string {
	var builder strings.Builder
	builder.WriteString("Deliver(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Delivers is a parsable slice of Deliver.
type Delivers []*Deliver

func (d Delivers) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
