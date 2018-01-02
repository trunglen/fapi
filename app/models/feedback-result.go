package models

import (
	mongodb "revel-modules/mongodb/app"
)

type FeedbackResult struct {
	BaseModel     `bson:",inline"`
	StoreID       string `bson:"store_id" json:"store_id"`
	ServiceID     string `bson:"service_id" json:"service_id"`
	TicketID      string `bson:"ticket_id" json:"ticket_id"`
	KioskID       string `bson:"kiosk_id" json:"kiosk_id"`
	CounterID     string `bson:"counter_id" json:"counter_id"`
	UserID        string `bson:"user_id" json:"user_id"`
	TransactionID string `bson:"transaction_id" json:"transaction_id"`
	Point         string `bson:"point" json:"point"`
	FeedbackID    string `bson:"feedback_id" json:"feedback_id"`
}

func newFeedbackResultCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("feedback_result")
}
func (f *FeedbackResult) Create() error {
	return newFeedbackResultCollection().Session.Insert(f)
}
