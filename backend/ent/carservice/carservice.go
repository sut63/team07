// Code generated by entc, DO NOT EDIT.

package carservice

const (
	// Label holds the string label denoting the carservice type in the database.
	Label = "carservice"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCustomer holds the string denoting the customer field in the database.
	FieldCustomer = "customer"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldServiceinfo holds the string denoting the serviceinfo field in the database.
	FieldServiceinfo = "serviceinfo"
	// FieldDatetime holds the string denoting the datetime field in the database.
	FieldDatetime = "datetime"

	// EdgeUserid holds the string denoting the userid edge name in mutations.
	EdgeUserid = "userid"
	// EdgeDisid holds the string denoting the disid edge name in mutations.
	EdgeDisid = "disid"
	// EdgeUrgentid holds the string denoting the urgentid edge name in mutations.
	EdgeUrgentid = "urgentid"

	// Table holds the table name of the carservice in the database.
	Table = "carservices"
	// UseridTable is the table the holds the userid relation/edge.
	UseridTable = "carservices"
	// UseridInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UseridInverseTable = "users"
	// UseridColumn is the table column denoting the userid relation/edge.
	UseridColumn = "user_id"
	// DisidTable is the table the holds the disid relation/edge.
	DisidTable = "carservices"
	// DisidInverseTable is the table name for the Distance entity.
	// It exists in this package in order to avoid circular dependency with the "distance" package.
	DisidInverseTable = "distances"
	// DisidColumn is the table column denoting the disid relation/edge.
	DisidColumn = "distance_disid"
	// UrgentidTable is the table the holds the urgentid relation/edge.
	UrgentidTable = "carservices"
	// UrgentidInverseTable is the table name for the Urgent entity.
	// It exists in this package in order to avoid circular dependency with the "urgent" package.
	UrgentidInverseTable = "urgents"
	// UrgentidColumn is the table column denoting the urgentid relation/edge.
	UrgentidColumn = "urgent_urgentid"
)

// Columns holds all SQL columns for carservice fields.
var Columns = []string{
	FieldID,
	FieldCustomer,
	FieldAge,
	FieldLocation,
	FieldServiceinfo,
	FieldDatetime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Carservice type.
var ForeignKeys = []string{
	"distance_disid",
	"urgent_urgentid",
	"user_id",
}

var (
	// CustomerValidator is a validator for the "customer" field. It is called by the builders before save.
	CustomerValidator func(string) error
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
	// LocationValidator is a validator for the "location" field. It is called by the builders before save.
	LocationValidator func(string) error
	// ServiceinfoValidator is a validator for the "serviceinfo" field. It is called by the builders before save.
	ServiceinfoValidator func(string) error
)
