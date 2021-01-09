// Code generated by entc, DO NOT EDIT.

package receive

const (
	// Label holds the string label denoting the receive type in the database.
	Label = "receive"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldReceivename holds the string denoting the receivename field in the database.
	FieldReceivename = "receivename"

	// EdgeReceiveid holds the string denoting the receiveid edge name in mutations.
	EdgeReceiveid = "receiveid"

	// Table holds the table name of the receive in the database.
	Table = "receives"
	// ReceiveidTable is the table the holds the receiveid relation/edge.
	ReceiveidTable = "transports"
	// ReceiveidInverseTable is the table name for the Transport entity.
	// It exists in this package in order to avoid circular dependency with the "transport" package.
	ReceiveidInverseTable = "transports"
	// ReceiveidColumn is the table column denoting the receiveid relation/edge.
	ReceiveidColumn = "receiveid"
)

// Columns holds all SQL columns for receive fields.
var Columns = []string{
	FieldID,
	FieldReceivename,
}
