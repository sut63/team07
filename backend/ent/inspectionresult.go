// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/jobposition"
)

// InspectionResult is the model entity for the InspectionResult schema.
type InspectionResult struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ResultName holds the value of the "result_name" field.
	ResultName string `json:"result_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InspectionResultQuery when eager-loading is set.
	Edges          InspectionResultEdges `json:"edges"`
	jobposition_id *int
}

// InspectionResultEdges holds the relations/edges for other nodes in the graph.
type InspectionResultEdges struct {
	// Carinspections holds the value of the carinspections edge.
	Carinspections []*CarInspection
	// Statusof holds the value of the statusof edge.
	Statusof []*Ambulance
	// Jobposition holds the value of the jobposition edge.
	Jobposition *JobPosition
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CarinspectionsOrErr returns the Carinspections value or an error if the edge
// was not loaded in eager-loading.
func (e InspectionResultEdges) CarinspectionsOrErr() ([]*CarInspection, error) {
	if e.loadedTypes[0] {
		return e.Carinspections, nil
	}
	return nil, &NotLoadedError{edge: "carinspections"}
}

// StatusofOrErr returns the Statusof value or an error if the edge
// was not loaded in eager-loading.
func (e InspectionResultEdges) StatusofOrErr() ([]*Ambulance, error) {
	if e.loadedTypes[1] {
		return e.Statusof, nil
	}
	return nil, &NotLoadedError{edge: "statusof"}
}

// JobpositionOrErr returns the Jobposition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InspectionResultEdges) JobpositionOrErr() (*JobPosition, error) {
	if e.loadedTypes[2] {
		if e.Jobposition == nil {
			// The edge jobposition was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: jobposition.Label}
		}
		return e.Jobposition, nil
	}
	return nil, &NotLoadedError{edge: "jobposition"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InspectionResult) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // result_name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*InspectionResult) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // jobposition_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InspectionResult fields.
func (ir *InspectionResult) assignValues(values ...interface{}) error {
	if m, n := len(values), len(inspectionresult.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ir.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field result_name", values[0])
	} else if value.Valid {
		ir.ResultName = value.String
	}
	values = values[1:]
	if len(values) == len(inspectionresult.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field jobposition_id", value)
		} else if value.Valid {
			ir.jobposition_id = new(int)
			*ir.jobposition_id = int(value.Int64)
		}
	}
	return nil
}

// QueryCarinspections queries the carinspections edge of the InspectionResult.
func (ir *InspectionResult) QueryCarinspections() *CarInspectionQuery {
	return (&InspectionResultClient{config: ir.config}).QueryCarinspections(ir)
}

// QueryStatusof queries the statusof edge of the InspectionResult.
func (ir *InspectionResult) QueryStatusof() *AmbulanceQuery {
	return (&InspectionResultClient{config: ir.config}).QueryStatusof(ir)
}

// QueryJobposition queries the jobposition edge of the InspectionResult.
func (ir *InspectionResult) QueryJobposition() *JobPositionQuery {
	return (&InspectionResultClient{config: ir.config}).QueryJobposition(ir)
}

// Update returns a builder for updating this InspectionResult.
// Note that, you need to call InspectionResult.Unwrap() before calling this method, if this InspectionResult
// was returned from a transaction, and the transaction was committed or rolled back.
func (ir *InspectionResult) Update() *InspectionResultUpdateOne {
	return (&InspectionResultClient{config: ir.config}).UpdateOne(ir)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ir *InspectionResult) Unwrap() *InspectionResult {
	tx, ok := ir.config.driver.(*txDriver)
	if !ok {
		panic("ent: InspectionResult is not a transactional entity")
	}
	ir.config.driver = tx.drv
	return ir
}

// String implements the fmt.Stringer.
func (ir *InspectionResult) String() string {
	var builder strings.Builder
	builder.WriteString("InspectionResult(")
	builder.WriteString(fmt.Sprintf("id=%v", ir.ID))
	builder.WriteString(", result_name=")
	builder.WriteString(ir.ResultName)
	builder.WriteByte(')')
	return builder.String()
}

// InspectionResults is a parsable slice of InspectionResult.
type InspectionResults []*InspectionResult

func (ir InspectionResults) config(cfg config) {
	for _i := range ir {
		ir[_i].config = cfg
	}
}
