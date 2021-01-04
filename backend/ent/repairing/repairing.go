// Code generated by entc, DO NOT EDIT.

package repairing

const (
	// Label holds the string label denoting the repairing type in the database.
	Label = "repairing"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRepairpart holds the string denoting the repairpart field in the database.
	FieldRepairpart = "repairpart"

	// EdgeRepairs holds the string denoting the repairs edge name in mutations.
	EdgeRepairs = "repairs"

	// Table holds the table name of the repairing in the database.
	Table = "repairings"
	// RepairsTable is the table the holds the repairs relation/edge.
	RepairsTable = "car_repairrecords"
	// RepairsInverseTable is the table name for the CarRepairrecord entity.
	// It exists in this package in order to avoid circular dependency with the "carrepairrecord" package.
	RepairsInverseTable = "car_repairrecords"
	// RepairsColumn is the table column denoting the repairs relation/edge.
	RepairsColumn = "repairing_id"
)

// Columns holds all SQL columns for repairing fields.
var Columns = []string{
	FieldID,
	FieldRepairpart,
}

var (
	// RepairpartValidator is a validator for the "repairpart" field. It is called by the builders before save.
	RepairpartValidator func(string) error
)
