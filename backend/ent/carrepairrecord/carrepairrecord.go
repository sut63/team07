// Code generated by entc, DO NOT EDIT.

package carrepairrecord

const (
	// Label holds the string label denoting the carrepairrecord type in the database.
	Label = "car_repairrecord"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDatetime holds the string denoting the datetime field in the database.
	FieldDatetime = "datetime"
	// FieldRepairdetail holds the string denoting the repairdetail field in the database.
	FieldRepairdetail = "repairdetail"
	// FieldRepaircost holds the string denoting the repaircost field in the database.
	FieldRepaircost = "repaircost"
	// FieldCarmaintenance holds the string denoting the carmaintenance field in the database.
	FieldCarmaintenance = "carmaintenance"

	// EdgeKeeper holds the string denoting the keeper edge name in mutations.
	EdgeKeeper = "keeper"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCarinspection holds the string denoting the carinspection edge name in mutations.
	EdgeCarinspection = "carinspection"

	// Table holds the table name of the carrepairrecord in the database.
	Table = "car_repairrecords"
	// KeeperTable is the table the holds the keeper relation/edge.
	KeeperTable = "car_repairrecords"
	// KeeperInverseTable is the table name for the Repairing entity.
	// It exists in this package in order to avoid circular dependency with the "repairing" package.
	KeeperInverseTable = "repairings"
	// KeeperColumn is the table column denoting the keeper relation/edge.
	KeeperColumn = "repairing_id"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "car_repairrecords"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// CarinspectionTable is the table the holds the carinspection relation/edge.
	CarinspectionTable = "car_repairrecords"
	// CarinspectionInverseTable is the table name for the CarInspection entity.
	// It exists in this package in order to avoid circular dependency with the "carinspection" package.
	CarinspectionInverseTable = "car_inspections"
	// CarinspectionColumn is the table column denoting the carinspection relation/edge.
	CarinspectionColumn = "carinspection_id"
)

// Columns holds all SQL columns for carrepairrecord fields.
var Columns = []string{
	FieldID,
	FieldDatetime,
	FieldRepairdetail,
	FieldRepaircost,
	FieldCarmaintenance,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the CarRepairrecord type.
var ForeignKeys = []string{
	"carinspection_id",
	"repairing_id",
	"user_id",
}

var (
	// RepairdetailValidator is a validator for the "repairdetail" field. It is called by the builders before save.
	RepairdetailValidator func(string) error
	// RepaircostValidator is a validator for the "repaircost" field. It is called by the builders before save.
	RepaircostValidator func(int) error
	// CarmaintenanceValidator is a validator for the "carmaintenance" field. It is called by the builders before save.
	CarmaintenanceValidator func(string) error
)
