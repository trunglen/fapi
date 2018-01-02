package models

type Staff struct {
	BaseModel `bson:",=inline"`
	Name      string `bson:"name" json:"name"`
	Store     string `bson:"store_id" json:"store_id"`
}
