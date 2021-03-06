// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/carbrand"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/insurance"
	"github.com/team07/app/ent/user"
)

// Ambulance is the model entity for the Ambulance schema.
type Ambulance struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Carregistration holds the value of the "carregistration" field.
	Carregistration string `json:"carregistration,omitempty"`
	// Enginepower holds the value of the "enginepower" field.
	Enginepower int `json:"enginepower,omitempty"`
	// Displacement holds the value of the "displacement" field.
	Displacement int `json:"displacement,omitempty"`
	// Registerat holds the value of the "registerat" field.
	Registerat time.Time `json:"registerat,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AmbulanceQuery when eager-loading is set.
	Edges        AmbulanceEdges `json:"edges"`
	brand_id     *int
	carstatus_id *int
	insurance_id *int
	user_id      *int
}

// AmbulanceEdges holds the relations/edges for other nodes in the graph.
type AmbulanceEdges struct {
	// Hasbrand holds the value of the hasbrand edge.
	Hasbrand *Carbrand
	// Hasinsurance holds the value of the hasinsurance edge.
	Hasinsurance *Insurance
	// Hasstatus holds the value of the hasstatus edge.
	Hasstatus *InspectionResult
	// Hasuser holds the value of the hasuser edge.
	Hasuser *User
	// Carinspections holds the value of the carinspections edge.
	Carinspections []*CarInspection
	// Carcheckinout holds the value of the carcheckinout edge.
	Carcheckinout []*CarCheckInOut
	// Ambulance holds the value of the ambulance edge.
	Ambulance []*Transport
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// HasbrandOrErr returns the Hasbrand value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AmbulanceEdges) HasbrandOrErr() (*Carbrand, error) {
	if e.loadedTypes[0] {
		if e.Hasbrand == nil {
			// The edge hasbrand was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: carbrand.Label}
		}
		return e.Hasbrand, nil
	}
	return nil, &NotLoadedError{edge: "hasbrand"}
}

// HasinsuranceOrErr returns the Hasinsurance value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AmbulanceEdges) HasinsuranceOrErr() (*Insurance, error) {
	if e.loadedTypes[1] {
		if e.Hasinsurance == nil {
			// The edge hasinsurance was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: insurance.Label}
		}
		return e.Hasinsurance, nil
	}
	return nil, &NotLoadedError{edge: "hasinsurance"}
}

// HasstatusOrErr returns the Hasstatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AmbulanceEdges) HasstatusOrErr() (*InspectionResult, error) {
	if e.loadedTypes[2] {
		if e.Hasstatus == nil {
			// The edge hasstatus was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: inspectionresult.Label}
		}
		return e.Hasstatus, nil
	}
	return nil, &NotLoadedError{edge: "hasstatus"}
}

// HasuserOrErr returns the Hasuser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AmbulanceEdges) HasuserOrErr() (*User, error) {
	if e.loadedTypes[3] {
		if e.Hasuser == nil {
			// The edge hasuser was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Hasuser, nil
	}
	return nil, &NotLoadedError{edge: "hasuser"}
}

// CarinspectionsOrErr returns the Carinspections value or an error if the edge
// was not loaded in eager-loading.
func (e AmbulanceEdges) CarinspectionsOrErr() ([]*CarInspection, error) {
	if e.loadedTypes[4] {
		return e.Carinspections, nil
	}
	return nil, &NotLoadedError{edge: "carinspections"}
}

// CarcheckinoutOrErr returns the Carcheckinout value or an error if the edge
// was not loaded in eager-loading.
func (e AmbulanceEdges) CarcheckinoutOrErr() ([]*CarCheckInOut, error) {
	if e.loadedTypes[5] {
		return e.Carcheckinout, nil
	}
	return nil, &NotLoadedError{edge: "carcheckinout"}
}

// AmbulanceOrErr returns the Ambulance value or an error if the edge
// was not loaded in eager-loading.
func (e AmbulanceEdges) AmbulanceOrErr() ([]*Transport, error) {
	if e.loadedTypes[6] {
		return e.Ambulance, nil
	}
	return nil, &NotLoadedError{edge: "ambulance"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ambulance) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // carregistration
		&sql.NullInt64{},  // enginepower
		&sql.NullInt64{},  // displacement
		&sql.NullTime{},   // registerat
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Ambulance) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // brand_id
		&sql.NullInt64{}, // carstatus_id
		&sql.NullInt64{}, // insurance_id
		&sql.NullInt64{}, // user_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ambulance fields.
func (a *Ambulance) assignValues(values ...interface{}) error {
	if m, n := len(values), len(ambulance.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field carregistration", values[0])
	} else if value.Valid {
		a.Carregistration = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field enginepower", values[1])
	} else if value.Valid {
		a.Enginepower = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field displacement", values[2])
	} else if value.Valid {
		a.Displacement = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field registerat", values[3])
	} else if value.Valid {
		a.Registerat = value.Time
	}
	values = values[4:]
	if len(values) == len(ambulance.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field brand_id", value)
		} else if value.Valid {
			a.brand_id = new(int)
			*a.brand_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field carstatus_id", value)
		} else if value.Valid {
			a.carstatus_id = new(int)
			*a.carstatus_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field insurance_id", value)
		} else if value.Valid {
			a.insurance_id = new(int)
			*a.insurance_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_id", value)
		} else if value.Valid {
			a.user_id = new(int)
			*a.user_id = int(value.Int64)
		}
	}
	return nil
}

// QueryHasbrand queries the hasbrand edge of the Ambulance.
func (a *Ambulance) QueryHasbrand() *CarbrandQuery {
	return (&AmbulanceClient{config: a.config}).QueryHasbrand(a)
}

// QueryHasinsurance queries the hasinsurance edge of the Ambulance.
func (a *Ambulance) QueryHasinsurance() *InsuranceQuery {
	return (&AmbulanceClient{config: a.config}).QueryHasinsurance(a)
}

// QueryHasstatus queries the hasstatus edge of the Ambulance.
func (a *Ambulance) QueryHasstatus() *InspectionResultQuery {
	return (&AmbulanceClient{config: a.config}).QueryHasstatus(a)
}

// QueryHasuser queries the hasuser edge of the Ambulance.
func (a *Ambulance) QueryHasuser() *UserQuery {
	return (&AmbulanceClient{config: a.config}).QueryHasuser(a)
}

// QueryCarinspections queries the carinspections edge of the Ambulance.
func (a *Ambulance) QueryCarinspections() *CarInspectionQuery {
	return (&AmbulanceClient{config: a.config}).QueryCarinspections(a)
}

// QueryCarcheckinout queries the carcheckinout edge of the Ambulance.
func (a *Ambulance) QueryCarcheckinout() *CarCheckInOutQuery {
	return (&AmbulanceClient{config: a.config}).QueryCarcheckinout(a)
}

// QueryAmbulance queries the ambulance edge of the Ambulance.
func (a *Ambulance) QueryAmbulance() *TransportQuery {
	return (&AmbulanceClient{config: a.config}).QueryAmbulance(a)
}

// Update returns a builder for updating this Ambulance.
// Note that, you need to call Ambulance.Unwrap() before calling this method, if this Ambulance
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Ambulance) Update() *AmbulanceUpdateOne {
	return (&AmbulanceClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Ambulance) Unwrap() *Ambulance {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ambulance is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Ambulance) String() string {
	var builder strings.Builder
	builder.WriteString("Ambulance(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", carregistration=")
	builder.WriteString(a.Carregistration)
	builder.WriteString(", enginepower=")
	builder.WriteString(fmt.Sprintf("%v", a.Enginepower))
	builder.WriteString(", displacement=")
	builder.WriteString(fmt.Sprintf("%v", a.Displacement))
	builder.WriteString(", registerat=")
	builder.WriteString(a.Registerat.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Ambulances is a parsable slice of Ambulance.
type Ambulances []*Ambulance

func (a Ambulances) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
