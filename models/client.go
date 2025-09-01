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
