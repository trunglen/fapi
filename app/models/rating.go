package models

import "strconv"

type Rating struct {
	Rating5     int `bson:"rating5" json:"rating5"`
	Rating4     int `bson:"rating4" json:"rating4"`
	Rating3     int `bson:"rating3" json:"rating3"`
	Rating2     int `bson:"rating2" json:"rating2"`
	Rating1     int `bson:"rating1" json:"rating1"`
	ReviewCount int `bson:"review_count" json:"review_count"`
}
type Target string

const (
	CRITERION = Target("criterion")
	STAFF     = Target("staff")
)

type TargetRating struct {
	Target   Target `bson:"target" json:"target"`
	TargetID string `bson:"rating1" json:"rating1"`
	Rating   Rating `bson:"rating" json:"rating"`
}

func (r *Rating) UpdateRating(rate int) {
	switch rate {
	case 1:
		r.Rating1 = r.Rating1 + 1
	case 2:
		r.Rating2 = r.Rating2 + 1
	case 3:
		r.Rating3 = r.Rating3 + 1
	case 4:
		r.Rating4 = r.Rating4 + 1
	case 5:
		r.Rating5 = r.Rating5 + 1
	}
	r.ReviewCount = r.ReviewCount + 1
}

func (r *Rating) GetAvgRating() float32 {
	var amount = r.Rating1 + r.Rating2 + r.Rating3 + r.Rating4 + r.Rating5
	if amount == 0 {
		return 5
	}
	return float32(r.Rating1+r.Rating2*2+r.Rating3*3+r.Rating4*4+r.Rating5*5) / float32(amount)
}

func (r *Rating) ToStringArr() []string {
	var res = make([]string, 0)
	var temp = func(rate int) string {
		return strconv.Itoa(rate)
	}
	res = append(res, temp(r.Rating1), temp(r.Rating2), temp(r.Rating3), temp(r.Rating4), temp(r.Rating5))
	return res
}
