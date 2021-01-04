// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/distance"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/jobposition"
	"github.com/team07/app/ent/repairing"
	"github.com/team07/app/ent/schema"
	"github.com/team07/app/ent/urgent"
	"github.com/team07/app/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	ambulanceFields := schema.Ambulance{}.Fields()
	_ = ambulanceFields
	// ambulanceDescRegisterat is the schema descriptor for registerat field.
	ambulanceDescRegisterat := ambulanceFields[1].Descriptor()
	// ambulance.DefaultRegisterat holds the default value on creation for the registerat field.
	ambulance.DefaultRegisterat = ambulanceDescRegisterat.Default.(func() time.Time)
	distanceFields := schema.Distance{}.Fields()
	_ = distanceFields
	// distanceDescDistance is the schema descriptor for Distance field.
	distanceDescDistance := distanceFields[0].Descriptor()
	// distance.DistanceValidator is a validator for the "Distance" field. It is called by the builders before save.
	distance.DistanceValidator = distanceDescDistance.Validators[0].(func(string) error)
	inspectionresultFields := schema.InspectionResult{}.Fields()
	_ = inspectionresultFields
	// inspectionresultDescResultName is the schema descriptor for result_name field.
	inspectionresultDescResultName := inspectionresultFields[0].Descriptor()
	// inspectionresult.ResultNameValidator is a validator for the "result_name" field. It is called by the builders before save.
	inspectionresult.ResultNameValidator = inspectionresultDescResultName.Validators[0].(func(string) error)
	jobpositionFields := schema.JobPosition{}.Fields()
	_ = jobpositionFields
	// jobpositionDescPositionName is the schema descriptor for position_name field.
	jobpositionDescPositionName := jobpositionFields[0].Descriptor()
	// jobposition.PositionNameValidator is a validator for the "position_name" field. It is called by the builders before save.
	jobposition.PositionNameValidator = jobpositionDescPositionName.Validators[0].(func(string) error)
	repairingFields := schema.Repairing{}.Fields()
	_ = repairingFields
	// repairingDescRepairpart is the schema descriptor for repairpart field.
	repairingDescRepairpart := repairingFields[0].Descriptor()
	// repairing.RepairpartValidator is a validator for the "repairpart" field. It is called by the builders before save.
	repairing.RepairpartValidator = repairingDescRepairpart.Validators[0].(func(string) error)
	urgentFields := schema.Urgent{}.Fields()
	_ = urgentFields
	// urgentDescUrgent is the schema descriptor for urgent field.
	urgentDescUrgent := urgentFields[0].Descriptor()
	// urgent.UrgentValidator is a validator for the "urgent" field. It is called by the builders before save.
	urgent.UrgentValidator = urgentDescUrgent.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
}
