// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/carcheckinout"
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/carservice"
	"github.com/team07/app/ent/distance"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/jobposition"
	"github.com/team07/app/ent/repairing"
	"github.com/team07/app/ent/schema"
	"github.com/team07/app/ent/transport"
	"github.com/team07/app/ent/urgent"
	"github.com/team07/app/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	ambulanceFields := schema.Ambulance{}.Fields()
	_ = ambulanceFields
	// ambulanceDescCarregistration is the schema descriptor for carregistration field.
	ambulanceDescCarregistration := ambulanceFields[0].Descriptor()
	// ambulance.CarregistrationValidator is a validator for the "carregistration" field. It is called by the builders before save.
	ambulance.CarregistrationValidator = func() func(string) error {
		validators := ambulanceDescCarregistration.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(carregistration string) error {
			for _, fn := range fns {
				if err := fn(carregistration); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// ambulanceDescEnginepower is the schema descriptor for enginepower field.
	ambulanceDescEnginepower := ambulanceFields[1].Descriptor()
	// ambulance.EnginepowerValidator is a validator for the "enginepower" field. It is called by the builders before save.
	ambulance.EnginepowerValidator = ambulanceDescEnginepower.Validators[0].(func(int) error)
	// ambulanceDescDisplacement is the schema descriptor for displacement field.
	ambulanceDescDisplacement := ambulanceFields[2].Descriptor()
	// ambulance.DisplacementValidator is a validator for the "displacement" field. It is called by the builders before save.
	ambulance.DisplacementValidator = ambulanceDescDisplacement.Validators[0].(func(int) error)
	carcheckinoutFields := schema.CarCheckInOut{}.Fields()
	_ = carcheckinoutFields
	// carcheckinoutDescPlace is the schema descriptor for place field.
	carcheckinoutDescPlace := carcheckinoutFields[1].Descriptor()
	// carcheckinout.PlaceValidator is a validator for the "place" field. It is called by the builders before save.
	carcheckinout.PlaceValidator = carcheckinoutDescPlace.Validators[0].(func(string) error)
	// carcheckinoutDescPerson is the schema descriptor for person field.
	carcheckinoutDescPerson := carcheckinoutFields[2].Descriptor()
	// carcheckinout.PersonValidator is a validator for the "person" field. It is called by the builders before save.
	carcheckinout.PersonValidator = func() func(int) error {
		validators := carcheckinoutDescPerson.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(person int) error {
			for _, fn := range fns {
				if err := fn(person); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// carcheckinoutDescDistance is the schema descriptor for distance field.
	carcheckinoutDescDistance := carcheckinoutFields[3].Descriptor()
	// carcheckinout.DistanceValidator is a validator for the "distance" field. It is called by the builders before save.
	carcheckinout.DistanceValidator = carcheckinoutDescDistance.Validators[0].(func(float64) error)
	// carcheckinoutDescCheckIn is the schema descriptor for checkIn field.
	carcheckinoutDescCheckIn := carcheckinoutFields[4].Descriptor()
	// carcheckinout.DefaultCheckIn holds the default value on creation for the checkIn field.
	carcheckinout.DefaultCheckIn = carcheckinoutDescCheckIn.Default.(func() time.Time)
	carinspectionFields := schema.CarInspection{}.Fields()
	_ = carinspectionFields
	// carinspectionDescWheelCenter is the schema descriptor for wheel_center field.
	carinspectionDescWheelCenter := carinspectionFields[0].Descriptor()
	// carinspection.WheelCenterValidator is a validator for the "wheel_center" field. It is called by the builders before save.
	carinspection.WheelCenterValidator = carinspectionDescWheelCenter.Validators[0].(func(float64) error)
	// carinspectionDescSoundLevel is the schema descriptor for sound_level field.
	carinspectionDescSoundLevel := carinspectionFields[1].Descriptor()
	// carinspection.SoundLevelValidator is a validator for the "sound_level" field. It is called by the builders before save.
	carinspection.SoundLevelValidator = carinspectionDescSoundLevel.Validators[0].(func(float64) error)
	// carinspectionDescBlacksmoke is the schema descriptor for blacksmoke field.
	carinspectionDescBlacksmoke := carinspectionFields[2].Descriptor()
	// carinspection.BlacksmokeValidator is a validator for the "blacksmoke" field. It is called by the builders before save.
	carinspection.BlacksmokeValidator = func() func(float64) error {
		validators := carinspectionDescBlacksmoke.Validators
		fns := [...]func(float64) error{
			validators[0].(func(float64) error),
			validators[1].(func(float64) error),
		}
		return func(blacksmoke float64) error {
			for _, fn := range fns {
				if err := fn(blacksmoke); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	carserviceFields := schema.Carservice{}.Fields()
	_ = carserviceFields
	// carserviceDescCustomer is the schema descriptor for customer field.
	carserviceDescCustomer := carserviceFields[0].Descriptor()
	// carservice.CustomerValidator is a validator for the "customer" field. It is called by the builders before save.
	carservice.CustomerValidator = carserviceDescCustomer.Validators[0].(func(string) error)
	// carserviceDescAge is the schema descriptor for age field.
	carserviceDescAge := carserviceFields[1].Descriptor()
	// carservice.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	carservice.AgeValidator = carserviceDescAge.Validators[0].(func(int) error)
	// carserviceDescLocation is the schema descriptor for location field.
	carserviceDescLocation := carserviceFields[2].Descriptor()
	// carservice.LocationValidator is a validator for the "location" field. It is called by the builders before save.
	carservice.LocationValidator = carserviceDescLocation.Validators[0].(func(string) error)
	// carserviceDescInformation is the schema descriptor for information field.
	carserviceDescInformation := carserviceFields[3].Descriptor()
	// carservice.InformationValidator is a validator for the "information" field. It is called by the builders before save.
	carservice.InformationValidator = carserviceDescInformation.Validators[0].(func(string) error)
	distanceFields := schema.Distance{}.Fields()
	_ = distanceFields
	// distanceDescDistance is the schema descriptor for Distance field.
	distanceDescDistance := distanceFields[0].Descriptor()
	// distance.DistanceValidator is a validator for the "Distance" field. It is called by the builders before save.
	distance.DistanceValidator = distanceDescDistance.Validators[0].(func(string) error)
	inspectionresultFields := schema.InspectionResult{}.Fields()
	_ = inspectionresultFields
	// inspectionresultDescResultName is the schema descriptor for result_name field.
	inspectionresultDescResultName := inspectionresultFields[0].Descriptor()
	// inspectionresult.ResultNameValidator is a validator for the "result_name" field. It is called by the builders before save.
	inspectionresult.ResultNameValidator = inspectionresultDescResultName.Validators[0].(func(string) error)
	jobpositionFields := schema.JobPosition{}.Fields()
	_ = jobpositionFields
	// jobpositionDescPositionName is the schema descriptor for position_name field.
	jobpositionDescPositionName := jobpositionFields[0].Descriptor()
	// jobposition.PositionNameValidator is a validator for the "position_name" field. It is called by the builders before save.
	jobposition.PositionNameValidator = jobpositionDescPositionName.Validators[0].(func(string) error)
	repairingFields := schema.Repairing{}.Fields()
	_ = repairingFields
	// repairingDescRepairpart is the schema descriptor for repairpart field.
	repairingDescRepairpart := repairingFields[0].Descriptor()
	// repairing.RepairpartValidator is a validator for the "repairpart" field. It is called by the builders before save.
	repairing.RepairpartValidator = repairingDescRepairpart.Validators[0].(func(string) error)
	transportFields := schema.Transport{}.Fields()
	_ = transportFields
	// transportDescSymptom is the schema descriptor for symptom field.
	transportDescSymptom := transportFields[0].Descriptor()
	// transport.SymptomValidator is a validator for the "symptom" field. It is called by the builders before save.
	transport.SymptomValidator = func() func(string) error {
		validators := transportDescSymptom.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(symptom string) error {
			for _, fn := range fns {
				if err := fn(symptom); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// transportDescDrugallergy is the schema descriptor for drugallergy field.
	transportDescDrugallergy := transportFields[1].Descriptor()
	// transport.DrugallergyValidator is a validator for the "drugallergy" field. It is called by the builders before save.
	transport.DrugallergyValidator = func() func(string) error {
		validators := transportDescDrugallergy.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(drugallergy string) error {
			for _, fn := range fns {
				if err := fn(drugallergy); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// transportDescNote is the schema descriptor for note field.
	transportDescNote := transportFields[2].Descriptor()
	// transport.NoteValidator is a validator for the "note" field. It is called by the builders before save.
	transport.NoteValidator = func() func(string) error {
		validators := transportDescNote.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(note string) error {
			for _, fn := range fns {
				if err := fn(note); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	urgentFields := schema.Urgent{}.Fields()
	_ = urgentFields
	// urgentDescUrgent is the schema descriptor for urgent field.
	urgentDescUrgent := urgentFields[0].Descriptor()
	// urgent.UrgentValidator is a validator for the "urgent" field. It is called by the builders before save.
	urgent.UrgentValidator = urgentDescUrgent.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
}
