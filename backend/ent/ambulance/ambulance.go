// Code generated by entc, DO NOT EDIT.

package ambulance

import (
	"time"
)

const (
	// Label holds the string label denoting the ambulance type in the database.
	Label = "ambulance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCarregistration holds the string denoting the carregistration field in the database.
	FieldCarregistration = "carregistration"
	// FieldRegisterAt holds the string denoting the register_at field in the database.
	FieldRegisterAt = "register_at"

	// EdgeHasbrand holds the string denoting the hasbrand edge name in mutations.
	EdgeHasbrand = "hasbrand"
	// EdgeHasinsurance holds the string denoting the hasinsurance edge name in mutations.
	EdgeHasinsurance = "hasinsurance"
	// EdgeHasstatus holds the string denoting the hasstatus edge name in mutations.
	EdgeHasstatus = "hasstatus"
	// EdgeHasuser holds the string denoting the hasuser edge name in mutations.
	EdgeHasuser = "hasuser"
	// EdgeCarinspections holds the string denoting the carinspections edge name in mutations.
	EdgeCarinspections = "carinspections"

	// Table holds the table name of the ambulance in the database.
	Table = "ambulances"
	// HasbrandTable is the table the holds the hasbrand relation/edge.
	HasbrandTable = "ambulances"
	// HasbrandInverseTable is the table name for the Carbrand entity.
	// It exists in this package in order to avoid circular dependency with the "carbrand" package.
	HasbrandInverseTable = "carbrands"
	// HasbrandColumn is the table column denoting the hasbrand relation/edge.
	HasbrandColumn = "brand_id"
	// HasinsuranceTable is the table the holds the hasinsurance relation/edge.
	HasinsuranceTable = "ambulances"
	// HasinsuranceInverseTable is the table name for the Insurance entity.
	// It exists in this package in order to avoid circular dependency with the "insurance" package.
	HasinsuranceInverseTable = "insurances"
	// HasinsuranceColumn is the table column denoting the hasinsurance relation/edge.
	HasinsuranceColumn = "insurance_id"
	// HasstatusTable is the table the holds the hasstatus relation/edge.
	HasstatusTable = "ambulances"
	// HasstatusInverseTable is the table name for the Carstatus entity.
	// It exists in this package in order to avoid circular dependency with the "carstatus" package.
	HasstatusInverseTable = "carstatuses"
	// HasstatusColumn is the table column denoting the hasstatus relation/edge.
	HasstatusColumn = "status_id"
	// HasuserTable is the table the holds the hasuser relation/edge.
	HasuserTable = "ambulances"
	// HasuserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	HasuserInverseTable = "users"
	// HasuserColumn is the table column denoting the hasuser relation/edge.
	HasuserColumn = "user_id"
	// CarinspectionsTable is the table the holds the carinspections relation/edge.
	CarinspectionsTable = "car_inspections"
	// CarinspectionsInverseTable is the table name for the CarInspection entity.
	// It exists in this package in order to avoid circular dependency with the "carinspection" package.
	CarinspectionsInverseTable = "car_inspections"
	// CarinspectionsColumn is the table column denoting the carinspections relation/edge.
	CarinspectionsColumn = "ambulance_id"
)

// Columns holds all SQL columns for ambulance fields.
var Columns = []string{
	FieldID,
	FieldCarregistration,
	FieldRegisterAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Ambulance type.
var ForeignKeys = []string{
	"brand_id",
	"status_id",
	"insurance_id",
	"user_id",
}

var (
	// DefaultRegisterAt holds the default value on creation for the register_at field.
	DefaultRegisterAt func() time.Time
)
