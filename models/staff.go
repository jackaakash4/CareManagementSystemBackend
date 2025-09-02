package staff

import(
	"time"

)



type Certification struct{
	Name			string		`bson: "name" json: "name"`
	Number			string		`bson: "number" json: "number'`
	IssuedBy		string		`bson:"issued_by" json: "issued_by"`
	IssuedDate		time.Time	`bson: "issued_date" json: "issued_date"`
	ExpiryDate		time.Time	`bson: "expiry_date" json: "expiry_date"`
	Status			string		`bson: "status" json: "status"` //completed, expired or pending
	DocumentPath	string		`bson: "document_path" json: "document_path"`
}


type Training struct{
	Name			string		`bson: "name" json: "name"`
	Provider		string		`bson: "provider" json: "provider"`
	CompletedAt		time.Time	`bson: "completed_at" json: "completed_at"`
	ExpiryDate		*time.Time	`bson: "expiry_date, omitempty" json: "expiry_date"`
	CertificationPath	string	`bson: "certification_path" json: "certification_path"`
	Status:			string		`bson: "status" json: "status"` //completed, expired or pending
}


type Availability struct{
	DayOfWeek		int		`bson: "day_of_week" json: "day_of_week"`		//0 = sunday, 1= monday, ... so on
	StartTime		string	`bson: "start_time" json: "start_time"`			//"09:00"
	EndTime			string	`bson: "end_time" json: "end_time"`				//"12:00"
	IsAvailable		bool	`bson: "is_available" json: "is_available"`
}



type Staff struct{
	ID				primitive.ObjectID		`bson: "_id, omitempty"	json: "id"`

	//PersonalInformation
	FirstName		string					`bson: "first_name" json: "first_name"`
	LastName		string					`bson: "last_name" json: "last_name"`
	DateOfBirth		time.Time				`bson: "date_of_birth" json: "date_of_birth"`
	Phone			string					`bson: "phone" json: "phone"`
	Email			string					`bson: "email" json: "email"`
	Address			string					`bson: "address" json: "address"`

	//EmergencyContact
	EmergencyContactName 	string			`bson: "emergency_contact_name" json: "emergency_contact_name"`
	EmergencyContactPhone	string			`bson: "emergency_contact_phone" json: "emergency_contact_phone"`

	//EmplouymentDetail

	EmployeeID		string			`bson: "employee_id" json: "employee_id"`
	JobRole			string			`bson: "job_role" json: "job_role"`			//support worker, nurse, coordinator
	HourlyRate		float64			`bson: "hourly_rate" json: "hourly_rate"`
	StartDate		time.Time		`bson: "start_date" json: "start_date"`
	EndDate			*time.Time		`bson: "end_date, omitempty" json: "end_date"`


	//Availability and roster
	Availability	[]Availability	`bson: "availability" json: "availability"`

	//Training and certificate
	Training		[]Training		`bson: "training" json: "training"`
	Certifications 	[]Certification	`bson: "certifications" json: "certifications"`

	//Documents
	Documents		[]Document		`bson: "documents" json: "documents"`

	//AssignedClient
	AssignedClients	[]primitive.ObjectId	`bson:"assigmed_clients" json: "assigned_clients"`

	//Metadate
	IsActive		bool			`bson: "is_active" json: "is_active"`
	CreatedAt		time.Time		`bson: "created_at" json: "created_at"`
	CreatedBy		primitive.ObjectID `bson: "created_by" json: "created_by"`
	UpdatedAt		time.Time		`bson: "updated_at" json: "updated_at"`
}


























