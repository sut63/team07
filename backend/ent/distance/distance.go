// Code generated by entc, DO NOT EDIT.

package distance

const (
	// Label holds the string label denoting the distance type in the database.
	Label = "distance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDistance holds the string denoting the distance field in the database.
	FieldDistance = "distance"

	// EdgeDisid holds the string denoting the disid edge name in mutations.
	EdgeDisid = "disid"

	// Table holds the table name of the distance in the database.
	Table = "distances"
	// DisidTable is the table the holds the disid relation/edge.
	DisidTable = "carservices"
	// DisidInverseTable is the table name for the Carservice entity.
	// It exists in this package in order to avoid circular dependency with the "carservice" package.
	DisidInverseTable = "carservices"
	// DisidColumn is the table column denoting the disid relation/edge.
	DisidColumn = "distance_disid"
)

// Columns holds all SQL columns for distance fields.
var Columns = []string{
	FieldID,
	FieldDistance,
}

var (
	// DistanceValidator is a validator for the "Distance" field. It is called by the builders before save.
	DistanceValidator func(string) error
)
