package models

type GuestLedger struct {
	Email   string `json:"email,omitempty" bson:"email,omitempty"`
	Message string `json:"message" bson:"message,omitempty"`
}
