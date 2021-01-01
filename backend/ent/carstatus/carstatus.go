// Code generated by entc, DO NOT EDIT.

package carstatus

const (
	// Label holds the string label denoting the carstatus type in the database.
	Label = "carstatus"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCarstatus holds the string denoting the carstatus field in the database.
	FieldCarstatus = "carstatus"

	// EdgeStatusof holds the string denoting the statusof edge name in mutations.
	EdgeStatusof = "statusof"

	// Table holds the table name of the carstatus in the database.
	Table = "carstatuses"
	// StatusofTable is the table the holds the statusof relation/edge.
	StatusofTable = "ambulances"
	// StatusofInverseTable is the table name for the Ambulance entity.
	// It exists in this package in order to avoid circular dependency with the "ambulance" package.
	StatusofInverseTable = "ambulances"
	// StatusofColumn is the table column denoting the statusof relation/edge.
	StatusofColumn = "status_id"
)

// Columns holds all SQL columns for carstatus fields.
var Columns = []string{
	FieldID,
	FieldCarstatus,
}
