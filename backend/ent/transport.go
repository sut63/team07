// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/hospital"
	"github.com/team07/app/ent/transport"
	"github.com/team07/app/ent/user"
)

// Transport is the model entity for the Transport schema.
type Transport struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Symptom holds the value of the "symptom" field.
	Symptom string `json:"symptom,omitempty"`
	// Drugallergy holds the value of the "drugallergy" field.
	Drugallergy string `json:"drugallergy,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransportQuery when eager-loading is set.
	Edges     TransportEdges `json:"edges"`
	ambulance *int
	receive   *int
	send      *int
	user      *int
}

// TransportEdges holds the relations/edges for other nodes in the graph.
type TransportEdges struct {
	// Send holds the value of the send edge.
	Send *Hospital
	// Receive holds the value of the receive edge.
	Receive *Hospital
	// User holds the value of the user edge.
	User *User
	// Ambulance holds the value of the ambulance edge.
	Ambulance *Ambulance
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// SendOrErr returns the Send value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransportEdges) SendOrErr() (*Hospital, error) {
	if e.loadedTypes[0] {
		if e.Send == nil {
			// The edge send was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: hospital.Label}
		}
		return e.Send, nil
	}
	return nil, &NotLoadedError{edge: "send"}
}

// ReceiveOrErr returns the Receive value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransportEdges) ReceiveOrErr() (*Hospital, error) {
	if e.loadedTypes[1] {
		if e.Receive == nil {
			// The edge receive was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: hospital.Label}
		}
		return e.Receive, nil
	}
	return nil, &NotLoadedError{edge: "receive"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransportEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// AmbulanceOrErr returns the Ambulance value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransportEdges) AmbulanceOrErr() (*Ambulance, error) {
	if e.loadedTypes[3] {
		if e.Ambulance == nil {
			// The edge ambulance was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: ambulance.Label}
		}
		return e.Ambulance, nil
	}
	return nil, &NotLoadedError{edge: "ambulance"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transport) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // symptom
		&sql.NullString{}, // drugallergy
		&sql.NullString{}, // note
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Transport) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // ambulance
		&sql.NullInt64{}, // receive
		&sql.NullInt64{}, // send
		&sql.NullInt64{}, // user
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transport fields.
func (t *Transport) assignValues(values ...interface{}) error {
	if m, n := len(values), len(transport.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field symptom", values[0])
	} else if value.Valid {
		t.Symptom = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field drugallergy", values[1])
	} else if value.Valid {
		t.Drugallergy = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field note", values[2])
	} else if value.Valid {
		t.Note = value.String
	}
	values = values[3:]
	if len(values) == len(transport.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field ambulance", value)
		} else if value.Valid {
			t.ambulance = new(int)
			*t.ambulance = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field receive", value)
		} else if value.Valid {
			t.receive = new(int)
			*t.receive = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field send", value)
		} else if value.Valid {
			t.send = new(int)
			*t.send = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user", value)
		} else if value.Valid {
			t.user = new(int)
			*t.user = int(value.Int64)
		}
	}
	return nil
}

// QuerySend queries the send edge of the Transport.
func (t *Transport) QuerySend() *HospitalQuery {
	return (&TransportClient{config: t.config}).QuerySend(t)
}

// QueryReceive queries the receive edge of the Transport.
func (t *Transport) QueryReceive() *HospitalQuery {
	return (&TransportClient{config: t.config}).QueryReceive(t)
}

// QueryUser queries the user edge of the Transport.
func (t *Transport) QueryUser() *UserQuery {
	return (&TransportClient{config: t.config}).QueryUser(t)
}

// QueryAmbulance queries the ambulance edge of the Transport.
func (t *Transport) QueryAmbulance() *AmbulanceQuery {
	return (&TransportClient{config: t.config}).QueryAmbulance(t)
}

// Update returns a builder for updating this Transport.
// Note that, you need to call Transport.Unwrap() before calling this method, if this Transport
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transport) Update() *TransportUpdateOne {
	return (&TransportClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Transport) Unwrap() *Transport {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transport is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transport) String() string {
	var builder strings.Builder
	builder.WriteString("Transport(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", symptom=")
	builder.WriteString(t.Symptom)
	builder.WriteString(", drugallergy=")
	builder.WriteString(t.Drugallergy)
	builder.WriteString(", note=")
	builder.WriteString(t.Note)
	builder.WriteByte(')')
	return builder.String()
}

// Transports is a parsable slice of Transport.
type Transports []*Transport

func (t Transports) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
