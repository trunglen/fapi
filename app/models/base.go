package models

import (
	"time"
)

type BaseModel struct {
	ID    string `bson:"_id" json:"id"`
	Dtime int64  `bson:"dtime" json:"dtime"`
	Mtime int64  `bson:"mtime" json:"mtime"`
	Ctime int64  `bson:"ctime" json:"ctime"`
}

func (b *BaseModel) BeforeCreate() {
	b.Ctime = time.Now().Unix()
}
func (b *BaseModel) BeforeUpdate() {
	b.Mtime = time.Now().Unix()
}
func (b *BaseModel) BeforeDelete() {
	b.Dtime = time.Now().Unix()
}
