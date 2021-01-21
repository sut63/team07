// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// AmbulancesColumns holds the columns for the "ambulances" table.
	AmbulancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "carregistration", Type: field.TypeString, Unique: true},
		{Name: "registerat", Type: field.TypeTime},
		{Name: "brand_id", Type: field.TypeInt, Nullable: true},
		{Name: "carstatus_id", Type: field.TypeInt, Nullable: true},
		{Name: "insurance_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// AmbulancesTable holds the schema information for the "ambulances" table.
	AmbulancesTable = &schema.Table{
		Name:       "ambulances",
		Columns:    AmbulancesColumns,
		PrimaryKey: []*schema.Column{AmbulancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "ambulances_carbrands_brandof",
				Columns: []*schema.Column{AmbulancesColumns[3]},

				RefColumns: []*schema.Column{CarbrandsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "ambulances_inspection_results_statusof",
				Columns: []*schema.Column{AmbulancesColumns[4]},

				RefColumns: []*schema.Column{InspectionResultsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "ambulances_insurances_insuranceof",
				Columns: []*schema.Column{AmbulancesColumns[5]},

				RefColumns: []*schema.Column{InsurancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "ambulances_users_userof",
				Columns: []*schema.Column{AmbulancesColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CarCheckInOutsColumns holds the columns for the "car_check_in_outs" table.
	CarCheckInOutsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "note", Type: field.TypeString},
		{Name: "check_in", Type: field.TypeTime},
		{Name: "check_out", Type: field.TypeTime},
		{Name: "ambulance", Type: field.TypeInt, Nullable: true},
		{Name: "purpose_carcheckinout", Type: field.TypeInt, Nullable: true},
		{Name: "name", Type: field.TypeInt, Nullable: true},
	}
	// CarCheckInOutsTable holds the schema information for the "car_check_in_outs" table.
	CarCheckInOutsTable = &schema.Table{
		Name:       "car_check_in_outs",
		Columns:    CarCheckInOutsColumns,
		PrimaryKey: []*schema.Column{CarCheckInOutsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "car_check_in_outs_ambulances_carcheckinout",
				Columns: []*schema.Column{CarCheckInOutsColumns[4]},

				RefColumns: []*schema.Column{AmbulancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "car_check_in_outs_purposes_carcheckinout",
				Columns: []*schema.Column{CarCheckInOutsColumns[5]},

				RefColumns: []*schema.Column{PurposesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "car_check_in_outs_users_carcheckinout",
				Columns: []*schema.Column{CarCheckInOutsColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CarInspectionsColumns holds the columns for the "car_inspections" table.
	CarInspectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "wheel_center", Type: field.TypeFloat64},
		{Name: "sound_level", Type: field.TypeFloat64},
		{Name: "blacksmoke", Type: field.TypeFloat64},
		{Name: "datetime", Type: field.TypeTime},
		{Name: "note", Type: field.TypeString},
		{Name: "ambulance_id", Type: field.TypeInt, Nullable: true},
		{Name: "inspectionresult_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// CarInspectionsTable holds the schema information for the "car_inspections" table.
	CarInspectionsTable = &schema.Table{
		Name:       "car_inspections",
		Columns:    CarInspectionsColumns,
		PrimaryKey: []*schema.Column{CarInspectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "car_inspections_ambulances_carinspections",
				Columns: []*schema.Column{CarInspectionsColumns[6]},

				RefColumns: []*schema.Column{AmbulancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "car_inspections_inspection_results_carinspections",
				Columns: []*schema.Column{CarInspectionsColumns[7]},

				RefColumns: []*schema.Column{InspectionResultsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "car_inspections_users_carinspections",
				Columns: []*schema.Column{CarInspectionsColumns[8]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CarRepairrecordsColumns holds the columns for the "car_repairrecords" table.
	CarRepairrecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "datetime", Type: field.TypeTime},
		{Name: "carinspection_id", Type: field.TypeInt, Nullable: true},
		{Name: "repairing_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// CarRepairrecordsTable holds the schema information for the "car_repairrecords" table.
	CarRepairrecordsTable = &schema.Table{
		Name:       "car_repairrecords",
		Columns:    CarRepairrecordsColumns,
		PrimaryKey: []*schema.Column{CarRepairrecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "car_repairrecords_car_inspections_carrepairrecords",
				Columns: []*schema.Column{CarRepairrecordsColumns[2]},

				RefColumns: []*schema.Column{CarInspectionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "car_repairrecords_repairings_repairs",
				Columns: []*schema.Column{CarRepairrecordsColumns[3]},

				RefColumns: []*schema.Column{RepairingsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "car_repairrecords_users_carrepairrecords",
				Columns: []*schema.Column{CarRepairrecordsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CarbrandsColumns holds the columns for the "carbrands" table.
	CarbrandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "brand", Type: field.TypeString, Unique: true},
	}
	// CarbrandsTable holds the schema information for the "carbrands" table.
	CarbrandsTable = &schema.Table{
		Name:        "carbrands",
		Columns:     CarbrandsColumns,
		PrimaryKey:  []*schema.Column{CarbrandsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CarregistersColumns holds the columns for the "carregisters" table.
	CarregistersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CarregistersTable holds the schema information for the "carregisters" table.
	CarregistersTable = &schema.Table{
		Name:        "carregisters",
		Columns:     CarregistersColumns,
		PrimaryKey:  []*schema.Column{CarregistersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CarservicesColumns holds the columns for the "carservices" table.
	CarservicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "customer", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
		{Name: "location", Type: field.TypeString},
		{Name: "information", Type: field.TypeString},
		{Name: "datetime", Type: field.TypeTime},
		{Name: "distance_disid", Type: field.TypeInt, Nullable: true},
		{Name: "urgent_urgentid", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// CarservicesTable holds the schema information for the "carservices" table.
	CarservicesTable = &schema.Table{
		Name:       "carservices",
		Columns:    CarservicesColumns,
		PrimaryKey: []*schema.Column{CarservicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "carservices_distances_disid",
				Columns: []*schema.Column{CarservicesColumns[6]},

				RefColumns: []*schema.Column{DistancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "carservices_urgents_urgentid",
				Columns: []*schema.Column{CarservicesColumns[7]},

				RefColumns: []*schema.Column{UrgentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "carservices_users_userid",
				Columns: []*schema.Column{CarservicesColumns[8]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DeliversColumns holds the columns for the "delivers" table.
	DeliversColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// DeliversTable holds the schema information for the "delivers" table.
	DeliversTable = &schema.Table{
		Name:        "delivers",
		Columns:     DeliversColumns,
		PrimaryKey:  []*schema.Column{DeliversColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DistancesColumns holds the columns for the "distances" table.
	DistancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "distance", Type: field.TypeString, Unique: true},
	}
	// DistancesTable holds the schema information for the "distances" table.
	DistancesTable = &schema.Table{
		Name:        "distances",
		Columns:     DistancesColumns,
		PrimaryKey:  []*schema.Column{DistancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// HospitalsColumns holds the columns for the "hospitals" table.
	HospitalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hospital", Type: field.TypeString, Unique: true},
	}
	// HospitalsTable holds the schema information for the "hospitals" table.
	HospitalsTable = &schema.Table{
		Name:        "hospitals",
		Columns:     HospitalsColumns,
		PrimaryKey:  []*schema.Column{HospitalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// InspectionResultsColumns holds the columns for the "inspection_results" table.
	InspectionResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "result_name", Type: field.TypeString, Unique: true},
		{Name: "jobposition_id", Type: field.TypeInt, Nullable: true},
	}
	// InspectionResultsTable holds the schema information for the "inspection_results" table.
	InspectionResultsTable = &schema.Table{
		Name:       "inspection_results",
		Columns:    InspectionResultsColumns,
		PrimaryKey: []*schema.Column{InspectionResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "inspection_results_job_positions_inspectionresults",
				Columns: []*schema.Column{InspectionResultsColumns[2]},

				RefColumns: []*schema.Column{JobPositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InsurancesColumns holds the columns for the "insurances" table.
	InsurancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "company", Type: field.TypeString, Unique: true},
	}
	// InsurancesTable holds the schema information for the "insurances" table.
	InsurancesTable = &schema.Table{
		Name:        "insurances",
		Columns:     InsurancesColumns,
		PrimaryKey:  []*schema.Column{InsurancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// JobPositionsColumns holds the columns for the "job_positions" table.
	JobPositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "position_name", Type: field.TypeString, Unique: true},
	}
	// JobPositionsTable holds the schema information for the "job_positions" table.
	JobPositionsTable = &schema.Table{
		Name:        "job_positions",
		Columns:     JobPositionsColumns,
		PrimaryKey:  []*schema.Column{JobPositionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PurposesColumns holds the columns for the "purposes" table.
	PurposesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "objective", Type: field.TypeString, Unique: true},
	}
	// PurposesTable holds the schema information for the "purposes" table.
	PurposesTable = &schema.Table{
		Name:        "purposes",
		Columns:     PurposesColumns,
		PrimaryKey:  []*schema.Column{PurposesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RepairingsColumns holds the columns for the "repairings" table.
	RepairingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "repairpart", Type: field.TypeString, Unique: true},
	}
	// RepairingsTable holds the schema information for the "repairings" table.
	RepairingsTable = &schema.Table{
		Name:        "repairings",
		Columns:     RepairingsColumns,
		PrimaryKey:  []*schema.Column{RepairingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TransportsColumns holds the columns for the "transports" table.
	TransportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "symptom", Type: field.TypeString},
		{Name: "drugallergy", Type: field.TypeString},
		{Name: "note", Type: field.TypeString},
		{Name: "ambulance", Type: field.TypeInt, Nullable: true},
		{Name: "receive", Type: field.TypeInt, Nullable: true},
		{Name: "send", Type: field.TypeInt, Nullable: true},
		{Name: "user", Type: field.TypeInt, Nullable: true},
	}
	// TransportsTable holds the schema information for the "transports" table.
	TransportsTable = &schema.Table{
		Name:       "transports",
		Columns:    TransportsColumns,
		PrimaryKey: []*schema.Column{TransportsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "transports_ambulances_ambulance",
				Columns: []*schema.Column{TransportsColumns[4]},

				RefColumns: []*schema.Column{AmbulancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "transports_hospitals_receive",
				Columns: []*schema.Column{TransportsColumns[5]},

				RefColumns: []*schema.Column{HospitalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "transports_hospitals_send",
				Columns: []*schema.Column{TransportsColumns[6]},

				RefColumns: []*schema.Column{HospitalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "transports_users_user",
				Columns: []*schema.Column{TransportsColumns[7]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UrgentsColumns holds the columns for the "urgents" table.
	UrgentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "urgent", Type: field.TypeString, Unique: true},
	}
	// UrgentsTable holds the schema information for the "urgents" table.
	UrgentsTable = &schema.Table{
		Name:        "urgents",
		Columns:     UrgentsColumns,
		PrimaryKey:  []*schema.Column{UrgentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "jobposition_id", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_job_positions_users",
				Columns: []*schema.Column{UsersColumns[4]},

				RefColumns: []*schema.Column{JobPositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AmbulancesTable,
		CarCheckInOutsTable,
		CarInspectionsTable,
		CarRepairrecordsTable,
		CarbrandsTable,
		CarregistersTable,
		CarservicesTable,
		DeliversTable,
		DistancesTable,
		HospitalsTable,
		InspectionResultsTable,
		InsurancesTable,
		JobPositionsTable,
		PurposesTable,
		RepairingsTable,
		TransportsTable,
		UrgentsTable,
		UsersTable,
	}
)

func init() {
	AmbulancesTable.ForeignKeys[0].RefTable = CarbrandsTable
	AmbulancesTable.ForeignKeys[1].RefTable = InspectionResultsTable
	AmbulancesTable.ForeignKeys[2].RefTable = InsurancesTable
	AmbulancesTable.ForeignKeys[3].RefTable = UsersTable
	CarCheckInOutsTable.ForeignKeys[0].RefTable = AmbulancesTable
	CarCheckInOutsTable.ForeignKeys[1].RefTable = PurposesTable
	CarCheckInOutsTable.ForeignKeys[2].RefTable = UsersTable
	CarInspectionsTable.ForeignKeys[0].RefTable = AmbulancesTable
	CarInspectionsTable.ForeignKeys[1].RefTable = InspectionResultsTable
	CarInspectionsTable.ForeignKeys[2].RefTable = UsersTable
	CarRepairrecordsTable.ForeignKeys[0].RefTable = CarInspectionsTable
	CarRepairrecordsTable.ForeignKeys[1].RefTable = RepairingsTable
	CarRepairrecordsTable.ForeignKeys[2].RefTable = UsersTable
	CarservicesTable.ForeignKeys[0].RefTable = DistancesTable
	CarservicesTable.ForeignKeys[1].RefTable = UrgentsTable
	CarservicesTable.ForeignKeys[2].RefTable = UsersTable
	InspectionResultsTable.ForeignKeys[0].RefTable = JobPositionsTable
	TransportsTable.ForeignKeys[0].RefTable = AmbulancesTable
	TransportsTable.ForeignKeys[1].RefTable = HospitalsTable
	TransportsTable.ForeignKeys[2].RefTable = HospitalsTable
	TransportsTable.ForeignKeys[3].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = JobPositionsTable
}
