package models

import (
	mongodb "revel-modules/mongodb/app"
)

type SurveyDevice struct {
	BaseModel `bson:",inline"`
	SurveyID  string   `bson:"survey_id" json:"survey_id"`
	DeviceID  []string `bson:"device_ids" json:"device_ids"`
}

func newSurveyDeviceCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("survey_device")
}

func (f *SurveyDevice) Create() error {
	return newSurveyDeviceCollection().Session.Insert(f)
}

func RemoveDevice() {

}
