// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// EdgeJobposition holds the string denoting the jobposition edge name in mutations.
	EdgeJobposition = "jobposition"
	// EdgeUserof holds the string denoting the userof edge name in mutations.
	EdgeUserof = "userof"
	// EdgeUserid holds the string denoting the userid edge name in mutations.
	EdgeUserid = "userid"
	// EdgeCarinspections holds the string denoting the carinspections edge name in mutations.
	EdgeCarinspections = "carinspections"
	// EdgeCarrepairrecords holds the string denoting the carrepairrecords edge name in mutations.
	EdgeCarrepairrecords = "carrepairrecords"
	// EdgeCarcheckinout holds the string denoting the carcheckinout edge name in mutations.
	EdgeCarcheckinout = "carcheckinout"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the user in the database.
	Table = "users"
	// JobpositionTable is the table the holds the jobposition relation/edge.
	JobpositionTable = "users"
	// JobpositionInverseTable is the table name for the JobPosition entity.
	// It exists in this package in order to avoid circular dependency with the "jobposition" package.
	JobpositionInverseTable = "job_positions"
	// JobpositionColumn is the table column denoting the jobposition relation/edge.
	JobpositionColumn = "jobposition_id"
	// UserofTable is the table the holds the userof relation/edge.
	UserofTable = "ambulances"
	// UserofInverseTable is the table name for the Ambulance entity.
	// It exists in this package in order to avoid circular dependency with the "ambulance" package.
	UserofInverseTable = "ambulances"
	// UserofColumn is the table column denoting the userof relation/edge.
	UserofColumn = "user_id"
	// UseridTable is the table the holds the userid relation/edge.
	UseridTable = "carservices"
	// UseridInverseTable is the table name for the Carservice entity.
	// It exists in this package in order to avoid circular dependency with the "carservice" package.
	UseridInverseTable = "carservices"
	// UseridColumn is the table column denoting the userid relation/edge.
	UseridColumn = "user_id"
	// CarinspectionsTable is the table the holds the carinspections relation/edge.
	CarinspectionsTable = "car_inspections"
	// CarinspectionsInverseTable is the table name for the CarInspection entity.
	// It exists in this package in order to avoid circular dependency with the "carinspection" package.
	CarinspectionsInverseTable = "car_inspections"
	// CarinspectionsColumn is the table column denoting the carinspections relation/edge.
	CarinspectionsColumn = "user_id"
	// CarrepairrecordsTable is the table the holds the carrepairrecords relation/edge.
	CarrepairrecordsTable = "car_repairrecords"
	// CarrepairrecordsInverseTable is the table name for the CarRepairrecord entity.
	// It exists in this package in order to avoid circular dependency with the "carrepairrecord" package.
	CarrepairrecordsInverseTable = "car_repairrecords"
	// CarrepairrecordsColumn is the table column denoting the carrepairrecords relation/edge.
	CarrepairrecordsColumn = "user_id"
	// CarcheckinoutTable is the table the holds the carcheckinout relation/edge.
	CarcheckinoutTable = "car_check_in_outs"
	// CarcheckinoutInverseTable is the table name for the CarCheckInOut entity.
	// It exists in this package in order to avoid circular dependency with the "carcheckinout" package.
	CarcheckinoutInverseTable = "car_check_in_outs"
	// CarcheckinoutColumn is the table column denoting the carcheckinout relation/edge.
	CarcheckinoutColumn = "name"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "transports"
	// UserInverseTable is the table name for the Transport entity.
	// It exists in this package in order to avoid circular dependency with the "transport" package.
	UserInverseTable = "transports"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"jobposition_id",
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
