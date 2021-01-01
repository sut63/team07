// Code generated by entc, DO NOT EDIT.

package jobposition

const (
	// Label holds the string label denoting the jobposition type in the database.
	Label = "job_position"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPositionName holds the string denoting the position_name field in the database.
	FieldPositionName = "position_name"

	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"

	// Table holds the table name of the jobposition in the database.
	Table = "job_positions"
	// UsersTable is the table the holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "jobposition_id"
)

// Columns holds all SQL columns for jobposition fields.
var Columns = []string{
	FieldID,
	FieldPositionName,
}

var (
	// PositionNameValidator is a validator for the "position_name" field. It is called by the builders before save.
	PositionNameValidator func(string) error
)
