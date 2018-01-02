package models

type Criterion struct {
	BaseModel `bson:",inline"`
	Name      string `bson:"name" json:"name"`
}
