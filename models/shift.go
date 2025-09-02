package shift

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shift struct {
	ID        primitive.ObjectID `bson: "_id, omitempty" json: "id"`
	ClientID  primitive.ObjectID `bson: "client_id" json: "client_id"`
	StaffID   primitive.ObjectID `bson: "staff_id" json: "staff_id"`
	StartTime time.Time          `bson: "start_time" json: "start_time"`
	EndTime   time.Time          `bson: "end_time" json: "end_time"`
	Status    string             `bson: "status" json: "status"` //scheduled, in_process, completed, cancelled

	//gps tracking(optional)
	CheckInLocation  *Location `bson: "checkin_location" json: "checkin_location"`
	CheckOutLocation *Location `bson: "checkout_location" json: "checkout_location"`

	//ShiftDetails
	ActualStartTime *time.Time `bson: "actual_start_time" json: "actual_start_time"`
	ActualEndTime   *time.Time `bson: "actual_end_time" json: "actual_end_time"`

	//Notes and Reports
	Notes []Note `bson: "notes" json: "notes"`

	CreatedAt time.Time `bson: "created_at" json: "created_at"`
	UpdatedAt time.Time `bson: "updated_at" json: "updated_at"`
}

type Location struct {
	Latitude  float64   `bson: "latitude" json: "latitude"`
	Longitude float64   `bson: "longitude" json: "longitude"`
	Address   string    `bson: "address" json: "address"`
	Timestamp time.Time `bson: "timestamp" json: "timestamp"`
}
