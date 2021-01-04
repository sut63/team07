// Code generated by entc, DO NOT EDIT.

package carinspection

const (
	// Label holds the string label denoting the carinspection type in the database.
	Label = "car_inspection"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDatetime holds the string denoting the datetime field in the database.
	FieldDatetime = "datetime"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeAmbulance holds the string denoting the ambulance edge name in mutations.
	EdgeAmbulance = "ambulance"
	// EdgeInspectionresult holds the string denoting the inspectionresult edge name in mutations.
	EdgeInspectionresult = "inspectionresult"
	// EdgeCarrepairrecords holds the string denoting the carrepairrecords edge name in mutations.
	EdgeCarrepairrecords = "carrepairrecords"

	// Table holds the table name of the carinspection in the database.
	Table = "car_inspections"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "car_inspections"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// AmbulanceTable is the table the holds the ambulance relation/edge.
	AmbulanceTable = "car_inspections"
	// AmbulanceInverseTable is the table name for the Ambulance entity.
	// It exists in this package in order to avoid circular dependency with the "ambulance" package.
	AmbulanceInverseTable = "ambulances"
	// AmbulanceColumn is the table column denoting the ambulance relation/edge.
	AmbulanceColumn = "ambulance_id"
	// InspectionresultTable is the table the holds the inspectionresult relation/edge.
	InspectionresultTable = "car_inspections"
	// InspectionresultInverseTable is the table name for the InspectionResult entity.
	// It exists in this package in order to avoid circular dependency with the "inspectionresult" package.
	InspectionresultInverseTable = "inspection_results"
	// InspectionresultColumn is the table column denoting the inspectionresult relation/edge.
	InspectionresultColumn = "inspectionresult_id"
	// CarrepairrecordsTable is the table the holds the carrepairrecords relation/edge.
	CarrepairrecordsTable = "car_repairrecords"
	// CarrepairrecordsInverseTable is the table name for the CarRepairrecord entity.
	// It exists in this package in order to avoid circular dependency with the "carrepairrecord" package.
	CarrepairrecordsInverseTable = "car_repairrecords"
	// CarrepairrecordsColumn is the table column denoting the carrepairrecords relation/edge.
	CarrepairrecordsColumn = "carinspection_id"
)

// Columns holds all SQL columns for carinspection fields.
var Columns = []string{
	FieldID,
	FieldDatetime,
	FieldNote,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the CarInspection type.
var ForeignKeys = []string{
	"ambulance_id",
	"inspectionresult_id",
	"user_id",
}
