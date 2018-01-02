package models

type Store struct {
	BaseModel `bson:",inline"`
	Name      string `bson:"name" json:"name"`
	Location  string `bson:"location" json:"location"`
	Level     int    `bson:"level" json:"level"`
}
