// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/carrepairrecord"
	"github.com/team07/app/ent/repairing"
	"github.com/team07/app/ent/user"
)

// CarRepairrecord is the model entity for the CarRepairrecord schema.
type CarRepairrecord struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime,omitempty"`
	// Partrepair holds the value of the "partrepair" field.
	Partrepair string `json:"partrepair,omitempty"`
	// Price holds the value of the "price" field.
	Price int `json:"price,omitempty"`
	// Techniciancomment holds the value of the "techniciancomment" field.
	Techniciancomment string `json:"techniciancomment,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CarRepairrecordQuery when eager-loading is set.
	Edges            CarRepairrecordEdges `json:"edges"`
	carinspection_id *int
	repairing_id     *int
	user_id          *int
}

// CarRepairrecordEdges holds the relations/edges for other nodes in the graph.
type CarRepairrecordEdges struct {
	// Keeper holds the value of the keeper edge.
	Keeper *Repairing
	// User holds the value of the user edge.
	User *User
	// Carinspection holds the value of the carinspection edge.
	Carinspection *CarInspection
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// KeeperOrErr returns the Keeper value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CarRepairrecordEdges) KeeperOrErr() (*Repairing, error) {
	if e.loadedTypes[0] {
		if e.Keeper == nil {
			// The edge keeper was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: repairing.Label}
		}
		return e.Keeper, nil
	}
	return nil, &NotLoadedError{edge: "keeper"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CarRepairrecordEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CarinspectionOrErr returns the Carinspection value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CarRepairrecordEdges) CarinspectionOrErr() (*CarInspection, error) {
	if e.loadedTypes[2] {
		if e.Carinspection == nil {
			// The edge carinspection was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: carinspection.Label}
		}
		return e.Carinspection, nil
	}
	return nil, &NotLoadedError{edge: "carinspection"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CarRepairrecord) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // datetime
		&sql.NullString{}, // partrepair
		&sql.NullInt64{},  // price
		&sql.NullString{}, // techniciancomment
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*CarRepairrecord) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // carinspection_id
		&sql.NullInt64{}, // repairing_id
		&sql.NullInt64{}, // user_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CarRepairrecord fields.
func (cr *CarRepairrecord) assignValues(values ...interface{}) error {
	if m, n := len(values), len(carrepairrecord.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	cr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field datetime", values[0])
	} else if value.Valid {
		cr.Datetime = value.Time
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field partrepair", values[1])
	} else if value.Valid {
		cr.Partrepair = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field price", values[2])
	} else if value.Valid {
		cr.Price = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field techniciancomment", values[3])
	} else if value.Valid {
		cr.Techniciancomment = value.String
	}
	values = values[4:]
	if len(values) == len(carrepairrecord.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field carinspection_id", value)
		} else if value.Valid {
			cr.carinspection_id = new(int)
			*cr.carinspection_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field repairing_id", value)
		} else if value.Valid {
			cr.repairing_id = new(int)
			*cr.repairing_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_id", value)
		} else if value.Valid {
			cr.user_id = new(int)
			*cr.user_id = int(value.Int64)
		}
	}
	return nil
}

// QueryKeeper queries the keeper edge of the CarRepairrecord.
func (cr *CarRepairrecord) QueryKeeper() *RepairingQuery {
	return (&CarRepairrecordClient{config: cr.config}).QueryKeeper(cr)
}

// QueryUser queries the user edge of the CarRepairrecord.
func (cr *CarRepairrecord) QueryUser() *UserQuery {
	return (&CarRepairrecordClient{config: cr.config}).QueryUser(cr)
}

// QueryCarinspection queries the carinspection edge of the CarRepairrecord.
func (cr *CarRepairrecord) QueryCarinspection() *CarInspectionQuery {
	return (&CarRepairrecordClient{config: cr.config}).QueryCarinspection(cr)
}

// Update returns a builder for updating this CarRepairrecord.
// Note that, you need to call CarRepairrecord.Unwrap() before calling this method, if this CarRepairrecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (cr *CarRepairrecord) Update() *CarRepairrecordUpdateOne {
	return (&CarRepairrecordClient{config: cr.config}).UpdateOne(cr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (cr *CarRepairrecord) Unwrap() *CarRepairrecord {
	tx, ok := cr.config.driver.(*txDriver)
	if !ok {
		panic("ent: CarRepairrecord is not a transactional entity")
	}
	cr.config.driver = tx.drv
	return cr
}

// String implements the fmt.Stringer.
func (cr *CarRepairrecord) String() string {
	var builder strings.Builder
	builder.WriteString("CarRepairrecord(")
	builder.WriteString(fmt.Sprintf("id=%v", cr.ID))
	builder.WriteString(", datetime=")
	builder.WriteString(cr.Datetime.Format(time.ANSIC))
	builder.WriteString(", partrepair=")
	builder.WriteString(cr.Partrepair)
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", cr.Price))
	builder.WriteString(", techniciancomment=")
	builder.WriteString(cr.Techniciancomment)
	builder.WriteByte(')')
	return builder.String()
}

// CarRepairrecords is a parsable slice of CarRepairrecord.
type CarRepairrecords []*CarRepairrecord

func (cr CarRepairrecords) config(cfg config) {
	for _i := range cr {
		cr[_i].config = cfg
	}
}
