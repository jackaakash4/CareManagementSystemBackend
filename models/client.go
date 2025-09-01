package client

import (
	"time",
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmergencyContact struct{
	Name 		    string 	    `bson: "name" json:"name"`
	Relationship	string	    `bson: "relationship" json: "relationship"`
	Phone		    string 	    `bson: "phone" json: "phone"`
	Email 		    string	    `bson: "email" json: "email"`
	IsPrimary	    bool	    `bson: "is_primary" json: "is_primary"`
}


type NextOfKin struct{
    Name            string      `bson: "name" json: "name"`
    Relationship    string      `bson: "relationship" json: "relationship"`
    Phone           string      `bson: "phone" json: "phone"`
    Email           string      `bson: "email" json: "email"`
    Address         string      `bson: "address" json: "email"`
    IsLegalGuardian string      `bson: "is_legel_guardian" json: "is_legel_guardian"`
}

type MedicalInfo struct{
    Allergies       []string    `bson: "allergies" json: "allergies"`
    Medications     []string    `bson: "medications" json: "Medications"`
    MedicalHistory  []string    `bson: "medical_history" json: "medical_history"`
    GP              string      `bson: "gp" json: "gp"`
    GPPhone         string      `bson: "gp_phone" json: "gp_phone"`
}


type CareInfo struct{
    CarePlan            string      `bson: "care_plan" json: "care_plan"`
    Goals               []string    `bson: "goals" json: "goals"`
    Supports            string      `bson: "supports" json: "supports"`
    MobilityRequirement string      `bson: "mobility_requirement" json: "mobility_requirement"`
    AccessibilityNeeds  string      `bson: "accessibility_needs" json: "accessibility_need"`
}

type RiskAssessment struct{
	Type			string		`bson:"type" json:"type"`
	Level			string		`bson:"level" json: "level"`
	Description		string		`bson:"description" json:"description"`
	Mitigation		string		`bson:"mitigation" json: "mitigation"`
	AssessedBy		string		`bson:"assessed_by" json: "assessed_by"`
	AssessedAt		time.Time	`bson:"assessed_at" json: "assessed_at"`
	ReviewDate		time.Time	`bson:"review_date" json: "review_date"`
}

type NDISInfo struct {
	NDISNumber			string		`bson:"ndis_number" json: "ndis_number"`
	PlanStartDate		time.Time	`bson:"plan_start_date"	json: "plan_start_date"`
	PlanEndDate			time.Time	`bson:"plan_end_date"	json: "plan_end_date"`
	ApprovedCategories	[]string	`bson:"approved_catogeries" json: "approved_categories"`
	ApprovedHours		int			`bson:"approved_hours" json: "approved_hours"`
	PlanManager			string		`bson: "plan_manager"	json: "plan_manager"`
	PlanManagerPhone	string		`bson: "plan_manager_phone" json: "plan_manager_phone"`
}

type Document struct{
	ID			primitive.ObjectID	`bson: "_id, omitempty" json: "id"`
	Name		string				`bson: "name" json: "name"`
	Type		string				`bson: "type" json: "type"`
	FilePath	string				`bson: "file_path" json: "file_path"`
	FileSize	int64				`bson: "file_size" json: "file_size"`
	MimeType	string				`bson: "mime_type" json: "mime_type"`
	UploadedBy	primitive.ObjectID	`bson:"uploaded_by" json: "uploaded_by"`
	UploadedAt	time.Time			`bson: "uploaded_at" json: "uploaded_at"`
	IsActive	bool				`bson: "is_active" json: "is_active"`
	}

type Note struct {
	ID			primitive.ObjectID	`bson: "_id, omitempty" json: "id"`
	Type		string				`bson: "type" json: "type"`
	Content		string				`bson: "content" json: "content"`
	CreatedBy	primitive.ObjectID	`bson: "created_by" json: "created_by"`
	CreatedAt	time.Time			`bson: "created_at" json: "created_at"`
	ShiftId		*primitive.ObjectID	`bson: "shift_id, omitempty" json: "shift_id"`
}


type Client struct{
	ID			primitive.ObjectID	`bson:"_id, omitempty" json: "id"`

	//Personal information

	FirstName	string				`bson:"first_name" json:"first_name"`
	LastName	string				`bson: "last_name" json: "last_name"`
	DateOfBirth	time.Time			`bson: "date_of_birth" json: "date_of_birth"`
	Gender		string				`bson: "gender" json: "gender"`
	Phone		string				`bson: "phone" json: "phone"`
	Email		string				`bson: "email" json: "email"`
	Address		string				`bson: "address" json: "address"`

	//Emergency contact
	EmergencyContacts	[]EmergencyContact	`bson: "emergency_contacts" json: "emergency_contacts"`
	NextOfKin			NextOfKin			`bson: "next_of_kin" json: "next_of_kin"`

	//Health and care information
	Medical				MedicalInfo			`bson: "medical" json: "medical"`
	Care				CareInfo			`bson: "care" json: "care"`
	RiskAssessments		[]RiskAssessment	`bson: "risk_assessements" json: "risk_assessments"`

	//NDIS information
	NDIS				NDISInfo			`bson: "ndis" json: "ndis"`

	//Documents and Notes
	Documents			[]Document			`bson: "documents" json: "documents"`
	Notes				[]Note				`bson: "notes" json: "notes"`

	//Metadata
	IsActive			bool				`bson: "is_active" json: "is_active"`
	CreatedAt			time.Time			`bson: "created_at" json: "created_at"`
	UpdatedAt			time.Time			`bson: "updated_at" json: "updated_at"`
	CreatedBy			primitive.ObjectID	`bson: "created_by" json: "updated_by"`
}


