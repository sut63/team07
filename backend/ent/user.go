// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team07/app/ent/jobposition"
	"github.com/team07/app/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges          UserEdges `json:"edges"`
	jobposition_id *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Jobposition holds the value of the jobposition edge.
	Jobposition *JobPosition
	// Userof holds the value of the userof edge.
	Userof []*Ambulance
	// Userid holds the value of the userid edge.
	Userid []*Carservice
	// Carinspections holds the value of the carinspections edge.
	Carinspections []*CarInspection
	// Carrepairrecords holds the value of the carrepairrecords edge.
	Carrepairrecords []*CarRepairrecord
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// JobpositionOrErr returns the Jobposition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) JobpositionOrErr() (*JobPosition, error) {
	if e.loadedTypes[0] {
		if e.Jobposition == nil {
			// The edge jobposition was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: jobposition.Label}
		}
		return e.Jobposition, nil
	}
	return nil, &NotLoadedError{edge: "jobposition"}
}

// UserofOrErr returns the Userof value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserofOrErr() ([]*Ambulance, error) {
	if e.loadedTypes[1] {
		return e.Userof, nil
	}
	return nil, &NotLoadedError{edge: "userof"}
}

// UseridOrErr returns the Userid value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UseridOrErr() ([]*Carservice, error) {
	if e.loadedTypes[2] {
		return e.Userid, nil
	}
	return nil, &NotLoadedError{edge: "userid"}
}

// CarinspectionsOrErr returns the Carinspections value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CarinspectionsOrErr() ([]*CarInspection, error) {
	if e.loadedTypes[3] {
		return e.Carinspections, nil
	}
	return nil, &NotLoadedError{edge: "carinspections"}
}

// CarrepairrecordsOrErr returns the Carrepairrecords value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CarrepairrecordsOrErr() ([]*CarRepairrecord, error) {
	if e.loadedTypes[4] {
		return e.Carrepairrecords, nil
	}
	return nil, &NotLoadedError{edge: "carrepairrecords"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // email
		&sql.NullString{}, // password
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // jobposition_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		u.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[1])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[2])
	} else if value.Valid {
		u.Password = value.String
	}
	values = values[3:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field jobposition_id", value)
		} else if value.Valid {
			u.jobposition_id = new(int)
			*u.jobposition_id = int(value.Int64)
		}
	}
	return nil
}

// QueryJobposition queries the jobposition edge of the User.
func (u *User) QueryJobposition() *JobPositionQuery {
	return (&UserClient{config: u.config}).QueryJobposition(u)
}

// QueryUserof queries the userof edge of the User.
func (u *User) QueryUserof() *AmbulanceQuery {
	return (&UserClient{config: u.config}).QueryUserof(u)
}

// QueryUserid queries the userid edge of the User.
func (u *User) QueryUserid() *CarserviceQuery {
	return (&UserClient{config: u.config}).QueryUserid(u)
}

// QueryCarinspections queries the carinspections edge of the User.
func (u *User) QueryCarinspections() *CarInspectionQuery {
	return (&UserClient{config: u.config}).QueryCarinspections(u)
}

// QueryCarrepairrecords queries the carrepairrecords edge of the User.
func (u *User) QueryCarrepairrecords() *CarRepairrecordQuery {
	return (&UserClient{config: u.config}).QueryCarrepairrecords(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
