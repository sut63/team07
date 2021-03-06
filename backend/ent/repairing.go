// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team07/app/ent/repairing"
)

// Repairing is the model entity for the Repairing schema.
type Repairing struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Repairpart holds the value of the "repairpart" field.
	Repairpart string `json:"repairpart,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RepairingQuery when eager-loading is set.
	Edges RepairingEdges `json:"edges"`
}

// RepairingEdges holds the relations/edges for other nodes in the graph.
type RepairingEdges struct {
	// Repairs holds the value of the repairs edge.
	Repairs []*CarRepairrecord
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RepairsOrErr returns the Repairs value or an error if the edge
// was not loaded in eager-loading.
func (e RepairingEdges) RepairsOrErr() ([]*CarRepairrecord, error) {
	if e.loadedTypes[0] {
		return e.Repairs, nil
	}
	return nil, &NotLoadedError{edge: "repairs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Repairing) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // repairpart
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Repairing fields.
func (r *Repairing) assignValues(values ...interface{}) error {
	if m, n := len(values), len(repairing.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field repairpart", values[0])
	} else if value.Valid {
		r.Repairpart = value.String
	}
	return nil
}

// QueryRepairs queries the repairs edge of the Repairing.
func (r *Repairing) QueryRepairs() *CarRepairrecordQuery {
	return (&RepairingClient{config: r.config}).QueryRepairs(r)
}

// Update returns a builder for updating this Repairing.
// Note that, you need to call Repairing.Unwrap() before calling this method, if this Repairing
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Repairing) Update() *RepairingUpdateOne {
	return (&RepairingClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Repairing) Unwrap() *Repairing {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Repairing is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Repairing) String() string {
	var builder strings.Builder
	builder.WriteString("Repairing(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", repairpart=")
	builder.WriteString(r.Repairpart)
	builder.WriteByte(')')
	return builder.String()
}

// Repairings is a parsable slice of Repairing.
type Repairings []*Repairing

func (r Repairings) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
