// Code generated by entc, DO NOT EDIT.

package hospital

const (
	// Label holds the string label denoting the hospital type in the database.
	Label = "hospital"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHospital holds the string denoting the hospital field in the database.
	FieldHospital = "hospital"

	// EdgeReceive holds the string denoting the receive edge name in mutations.
	EdgeReceive = "receive"
	// EdgeSend holds the string denoting the send edge name in mutations.
	EdgeSend = "send"

	// Table holds the table name of the hospital in the database.
	Table = "hospitals"
	// ReceiveTable is the table the holds the receive relation/edge.
	ReceiveTable = "transports"
	// ReceiveInverseTable is the table name for the Transport entity.
	// It exists in this package in order to avoid circular dependency with the "transport" package.
	ReceiveInverseTable = "transports"
	// ReceiveColumn is the table column denoting the receive relation/edge.
	ReceiveColumn = "receive"
	// SendTable is the table the holds the send relation/edge.
	SendTable = "transports"
	// SendInverseTable is the table name for the Transport entity.
	// It exists in this package in order to avoid circular dependency with the "transport" package.
	SendInverseTable = "transports"
	// SendColumn is the table column denoting the send relation/edge.
	SendColumn = "send"
)

// Columns holds all SQL columns for hospital fields.
var Columns = []string{
	FieldID,
	FieldHospital,
}
