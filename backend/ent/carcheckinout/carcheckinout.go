// Code generated by entc, DO NOT EDIT.

package carcheckinout

import (
	"time"
)

const (
	// Label holds the string label denoting the carcheckinout type in the database.
	Label = "car_check_in_out"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldPlace holds the string denoting the place field in the database.
	FieldPlace = "place"
	// FieldPerson holds the string denoting the person field in the database.
	FieldPerson = "person"
	// FieldDistance holds the string denoting the distance field in the database.
	FieldDistance = "distance"
	// FieldCheckIn holds the string denoting the checkin field in the database.
	FieldCheckIn = "check_in"
	// FieldCheckOut holds the string denoting the checkout field in the database.
	FieldCheckOut = "check_out"

	// EdgeAmbulance holds the string denoting the ambulance edge name in mutations.
	EdgeAmbulance = "ambulance"
	// EdgeName holds the string denoting the name edge name in mutations.
	EdgeName = "name"
	// EdgePurpose holds the string denoting the purpose edge name in mutations.
	EdgePurpose = "purpose"

	// Table holds the table name of the carcheckinout in the database.
	Table = "car_check_in_outs"
	// AmbulanceTable is the table the holds the ambulance relation/edge.
	AmbulanceTable = "car_check_in_outs"
	// AmbulanceInverseTable is the table name for the Ambulance entity.
	// It exists in this package in order to avoid circular dependency with the "ambulance" package.
	AmbulanceInverseTable = "ambulances"
	// AmbulanceColumn is the table column denoting the ambulance relation/edge.
	AmbulanceColumn = "ambulance"
	// NameTable is the table the holds the name relation/edge.
	NameTable = "car_check_in_outs"
	// NameInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	NameInverseTable = "users"
	// NameColumn is the table column denoting the name relation/edge.
	NameColumn = "name"
	// PurposeTable is the table the holds the purpose relation/edge.
	PurposeTable = "car_check_in_outs"
	// PurposeInverseTable is the table name for the Purpose entity.
	// It exists in this package in order to avoid circular dependency with the "purpose" package.
	PurposeInverseTable = "purposes"
	// PurposeColumn is the table column denoting the purpose relation/edge.
	PurposeColumn = "purpose_carcheckinout"
)

// Columns holds all SQL columns for carcheckinout fields.
var Columns = []string{
	FieldID,
	FieldNote,
	FieldPlace,
	FieldPerson,
	FieldDistance,
	FieldCheckIn,
	FieldCheckOut,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the CarCheckInOut type.
var ForeignKeys = []string{
	"ambulance",
	"purpose_carcheckinout",
	"name",
}

var (
	// PlaceValidator is a validator for the "place" field. It is called by the builders before save.
	PlaceValidator func(string) error
	// PersonValidator is a validator for the "person" field. It is called by the builders before save.
	PersonValidator func(int) error
	// DistanceValidator is a validator for the "distance" field. It is called by the builders before save.
	DistanceValidator func(float64) error
	// DefaultCheckIn holds the default value on creation for the checkIn field.
	DefaultCheckIn func() time.Time
)
