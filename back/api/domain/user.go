package domain

type User struct {
	ID            int
	AuthID        int
	Period        string
	FirstName     string
	LastName      string
	DepartmentID  int
	JobID         int
	EnrollmentFlg bool
	AdminFlg      bool
}

type Users []User
