// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team07/app/ent/carregister"
)

// Carregister is the model entity for the Carregister schema.
type Carregister struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Carregister) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Carregister fields.
func (c *Carregister) assignValues(values ...interface{}) error {
	if m, n := len(values), len(carregister.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	return nil
}

// Update returns a builder for updating this Carregister.
// Note that, you need to call Carregister.Unwrap() before calling this method, if this Carregister
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Carregister) Update() *CarregisterUpdateOne {
	return (&CarregisterClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Carregister) Unwrap() *Carregister {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Carregister is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Carregister) String() string {
	var builder strings.Builder
	builder.WriteString("Carregister(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Carregisters is a parsable slice of Carregister.
type Carregisters []*Carregister

func (c Carregisters) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
