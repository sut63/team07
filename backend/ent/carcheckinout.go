// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/carcheckinout"
	"github.com/team07/app/ent/purpose"
	"github.com/team07/app/ent/user"
)

// CarCheckInOut is the model entity for the CarCheckInOut schema.
type CarCheckInOut struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// Place holds the value of the "place" field.
	Place string `json:"place,omitempty"`
	// Person holds the value of the "person" field.
	Person int `json:"person,omitempty"`
	// Distance holds the value of the "distance" field.
	Distance float64 `json:"distance,omitempty"`
	// CheckIn holds the value of the "checkIn" field.
	CheckIn time.Time `json:"checkIn,omitempty"`
	// CheckOut holds the value of the "checkOut" field.
	CheckOut time.Time `json:"checkOut,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CarCheckInOutQuery when eager-loading is set.
	Edges                 CarCheckInOutEdges `json:"edges"`
	ambulance             *int
	purpose_carcheckinout *int
	name                  *int
}

// CarCheckInOutEdges holds the relations/edges for other nodes in the graph.
type CarCheckInOutEdges struct {
	// Ambulance holds the value of the ambulance edge.
	Ambulance *Ambulance
	// Name holds the value of the name edge.
	Name *User
	// Purpose holds the value of the purpose edge.
	Purpose *Purpose
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// AmbulanceOrErr returns the Ambulance value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CarCheckInOutEdges) AmbulanceOrErr() (*Ambulance, error) {
	if e.loadedTypes[0] {
		if e.Ambulance == nil {
			// The edge ambulance was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: ambulance.Label}
		}
		return e.Ambulance, nil
	}
	return nil, &NotLoadedError{edge: "ambulance"}
}

// NameOrErr returns the Name value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CarCheckInOutEdges) NameOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Name == nil {
			// The edge name was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Name, nil
	}
	return nil, &NotLoadedError{edge: "name"}
}

// PurposeOrErr returns the Purpose value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CarCheckInOutEdges) PurposeOrErr() (*Purpose, error) {
	if e.loadedTypes[2] {
		if e.Purpose == nil {
			// The edge purpose was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: purpose.Label}
		}
		return e.Purpose, nil
	}
	return nil, &NotLoadedError{edge: "purpose"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CarCheckInOut) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},   // id
		&sql.NullString{},  // note
		&sql.NullString{},  // place
		&sql.NullInt64{},   // person
		&sql.NullFloat64{}, // distance
		&sql.NullTime{},    // checkIn
		&sql.NullTime{},    // checkOut
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*CarCheckInOut) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // ambulance
		&sql.NullInt64{}, // purpose_carcheckinout
		&sql.NullInt64{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CarCheckInOut fields.
func (ccio *CarCheckInOut) assignValues(values ...interface{}) error {
	if m, n := len(values), len(carcheckinout.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ccio.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field note", values[0])
	} else if value.Valid {
		ccio.Note = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field place", values[1])
	} else if value.Valid {
		ccio.Place = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field person", values[2])
	} else if value.Valid {
		ccio.Person = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field distance", values[3])
	} else if value.Valid {
		ccio.Distance = value.Float64
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field checkIn", values[4])
	} else if value.Valid {
		ccio.CheckIn = value.Time
	}
	if value, ok := values[5].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field checkOut", values[5])
	} else if value.Valid {
		ccio.CheckOut = value.Time
	}
	values = values[6:]
	if len(values) == len(carcheckinout.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field ambulance", value)
		} else if value.Valid {
			ccio.ambulance = new(int)
			*ccio.ambulance = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field purpose_carcheckinout", value)
		} else if value.Valid {
			ccio.purpose_carcheckinout = new(int)
			*ccio.purpose_carcheckinout = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field name", value)
		} else if value.Valid {
			ccio.name = new(int)
			*ccio.name = int(value.Int64)
		}
	}
	return nil
}

// QueryAmbulance queries the ambulance edge of the CarCheckInOut.
func (ccio *CarCheckInOut) QueryAmbulance() *AmbulanceQuery {
	return (&CarCheckInOutClient{config: ccio.config}).QueryAmbulance(ccio)
}

// QueryName queries the name edge of the CarCheckInOut.
func (ccio *CarCheckInOut) QueryName() *UserQuery {
	return (&CarCheckInOutClient{config: ccio.config}).QueryName(ccio)
}

// QueryPurpose queries the purpose edge of the CarCheckInOut.
func (ccio *CarCheckInOut) QueryPurpose() *PurposeQuery {
	return (&CarCheckInOutClient{config: ccio.config}).QueryPurpose(ccio)
}

// Update returns a builder for updating this CarCheckInOut.
// Note that, you need to call CarCheckInOut.Unwrap() before calling this method, if this CarCheckInOut
// was returned from a transaction, and the transaction was committed or rolled back.
func (ccio *CarCheckInOut) Update() *CarCheckInOutUpdateOne {
	return (&CarCheckInOutClient{config: ccio.config}).UpdateOne(ccio)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ccio *CarCheckInOut) Unwrap() *CarCheckInOut {
	tx, ok := ccio.config.driver.(*txDriver)
	if !ok {
		panic("ent: CarCheckInOut is not a transactional entity")
	}
	ccio.config.driver = tx.drv
	return ccio
}

// String implements the fmt.Stringer.
func (ccio *CarCheckInOut) String() string {
	var builder strings.Builder
	builder.WriteString("CarCheckInOut(")
	builder.WriteString(fmt.Sprintf("id=%v", ccio.ID))
	builder.WriteString(", note=")
	builder.WriteString(ccio.Note)
	builder.WriteString(", place=")
	builder.WriteString(ccio.Place)
	builder.WriteString(", person=")
	builder.WriteString(fmt.Sprintf("%v", ccio.Person))
	builder.WriteString(", distance=")
	builder.WriteString(fmt.Sprintf("%v", ccio.Distance))
	builder.WriteString(", checkIn=")
	builder.WriteString(ccio.CheckIn.Format(time.ANSIC))
	builder.WriteString(", checkOut=")
	builder.WriteString(ccio.CheckOut.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CarCheckInOuts is a parsable slice of CarCheckInOut.
type CarCheckInOuts []*CarCheckInOut

func (ccio CarCheckInOuts) config(cfg config) {
	for _i := range ccio {
		ccio[_i].config = cfg
	}
}
