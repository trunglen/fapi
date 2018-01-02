package models

import (
	mongodb "github.com/trunglen/revel-modules/mongodb/app"

	"g/x/math"

	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type Smiley struct {
	BaseModel `bson:",inline"`
	Name      string `bson:"name" json:"name"`
	Point     int    `bson:"point" json:"point"`
}

func (s *Smiley) Validate(validation *revel.Validation) {
	validation.Check(s.Name, revel.Required{}).Message("Name required")
	validation.Check(s.Point, revel.Min{Min: 0}).Message("Point at least 0")
}
func newSmileyCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("smiley")
}
func GetSmileys() (error, []Smiley) {
	var smileys []Smiley
	err := newSmileyCollection().Session.Find(bson.M{}).All(&smileys)
	if err != nil {
		revel.AppLog.Errorf(err.Error())
	}
	return err, smileys
}
func (c *Smiley) Create() error {
	c.ID = math.RandString("smiley", 4)
	c.BeforeCreate()
	return newSmileyCollection().Session.Insert(c)
}
