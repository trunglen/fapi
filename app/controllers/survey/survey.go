package survey

import (
	"fapi/app/models"
	"g/x/web"
	xcontroller "revel-modules/xcontroller/app/controllers"

	"github.com/revel/revel"
)

type Survey struct {
	xcontroller.XController
}

func (c Survey) Add() revel.Result {
	var survey *models.Survey
	c.MustDecodeBody(&survey)
	survey.Validate(c.Validation)
	web.AssertValidation(c.Validation.Errors)
	web.AssertNil(survey.Create())
	return c.SendData(survey)
}
func (c Survey) List() revel.Result {
	surveys, err := models.ListSurvey()
	web.AssertNil(err)
	return c.SendData(surveys)
}
func (c Survey) Delete(id string) revel.Result {
	web.AssertNil(models.DeleteSurveyByID(id))
	return c.Success()
}

func (c Survey) AddSurveyDevice() revel.Result {
	var surveyDevice = struct {
		SurveyID string `bson:"survey_id" json:"survey_id"`
		DeviceID string `bson:"device_id" json:"device_id"`
	}{}
	c.MustDecodeBody(&surveyDevice)
	web.AssertNil(models.AddDeviceToSurvey(surveyDevice.DeviceID, surveyDevice.SurveyID))
	return c.Success()
}
