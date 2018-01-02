package models

import (
	"g/x/math"
	mongodb "revel-modules/mongodb/app"

	"gopkg.in/mgo.v2/bson"

	"github.com/revel/revel"
)

type Survey struct {
	BaseModel `bson:",inline"`
	Name      string     `bson:"name" json:"name"`
	Questions []Question `bson:"questions" json:"questions"`
	DeviceID  []string   `bson:"device_ids" json:"device_ids"`
}
type Question struct {
	ID      string     `bson:"id" json:"id"`
	Content string     `bson:"content" json:"content"`
	Summary string     `bson:"summary" json:"summary"`
	Type    SurveyType `bson:"type" json:"type"`
	Point   int        `bson:"point" json:"point"`
	Manded  bool       `bson:"manded" json:"manded"`
	Answers []Answer   `bson:"answers" json:"answers"`
}
type Answer struct {
	Icon    string `bson:"icon" json:"icon"`
	Content string `bson:"content" json:"content"`
	Point   int    `bson:"point" json:"point"`
	Link    string `bson:"link" json:"link"`
}
type SurveyType string

const (
	SINGLE   = SurveyType("single")
	MULTIPLE = SurveyType("multiple")
	ANSWER   = SurveyType("answer")
)

func (s *Survey) Validate(validation *revel.Validation) {
	validation.Check(s.Name, revel.Required{}).Message("Name required")
	validation.Check(s, QuestionValidator{})
}
func newSurveyCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("survey")
}

func DeleteSurveyByID(id string) error {
	return newSurveyCollection().Session.RemoveId(id)
}

func ListSurvey() ([]*Survey, error) {
	var surveys []*Survey
	err := newSurveyCollection().Session.Find(bson.M{}).All(&surveys)
	return surveys, err
}

func (s *Survey) Create() error {
	s.ID = math.RandString("srv", 4)
	s.BeforeCreate()
	return newSurveyCollection().Session.Insert(s)
}

func AddDeviceToSurvey(deviceID string, surveyID string) error {
	return newSurveyDeviceCollection().Session.UpdateId(surveyID, bson.M{
		"$addToSet": bson.M{
			"device_ids": []string{deviceID},
		},
	})
}

type QuestionValidator struct {
}

var questionError string

func (v QuestionValidator) IsSatisfied(s interface{}) bool {
	if survey, ok := s.(*Survey); ok {
		if len(survey.Questions) == 0 {
			questionError = "Question at least 1"
			return false
		}
		for _, q := range survey.Questions {
			if q.Content == "" {
				questionError = "Content is not emplty"
				return false
			}
			if q.Type != SINGLE && q.Type != MULTIPLE && q.Type != ANSWER {
				questionError = "Type is a value of {single,multiple,answer}"
				return false
			}
			if q.Point < 0 {
				questionError = "Point must be >0"
			}
			if len(q.Answers) == 0 {
				if q.Type != ANSWER {
					questionError = "Answer at least 1"
					return false
				}
			} else {
				for _, item := range q.Answers {
					if item.Content == "" {
						questionError = "Content is not emplty"
						return false
					}
					if q.Point < 0 {
						questionError = "Point must be >0"
					}
				}
			}
		}
		return true
	}
	return false
}

func (v QuestionValidator) DefaultMessage() string {
	return questionError
}
