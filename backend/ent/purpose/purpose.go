// Code generated by entc, DO NOT EDIT.

package purpose

const (
	// Label holds the string label denoting the purpose type in the database.
	Label = "purpose"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldObjective holds the string denoting the objective field in the database.
	FieldObjective = "objective"

	// EdgeCarcheckinout holds the string denoting the carcheckinout edge name in mutations.
	EdgeCarcheckinout = "carcheckinout"

	// Table holds the table name of the purpose in the database.
	Table = "purposes"
	// CarcheckinoutTable is the table the holds the carcheckinout relation/edge.
	CarcheckinoutTable = "car_check_in_outs"
	// CarcheckinoutInverseTable is the table name for the CarCheckInOut entity.
	// It exists in this package in order to avoid circular dependency with the "carcheckinout" package.
	CarcheckinoutInverseTable = "car_check_in_outs"
	// CarcheckinoutColumn is the table column denoting the carcheckinout relation/edge.
	CarcheckinoutColumn = "purpose_carcheckinout"
)

// Columns holds all SQL columns for purpose fields.
var Columns = []string{
	FieldID,
	FieldObjective,
}
