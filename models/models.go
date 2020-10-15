package models

// Webstruct contains Data that will be displayed on the website
type Webstruct struct {
	//CompanyName string
	AppSpec App
	User    User
	IsAuth  bool
	Message string
}

// User struct contains data that corresponds to a user
type User struct {
	Fname     string      `db:"fname"`
	Lname     string      `db:"lname"`
	UID       int         `db:"uid"`
	EmpStatus string      `db:"emp_status"`
	Joining   string      `db:"joining"`
	Role      Role        `db:"role"`
	Contact   Contact     `db:"contact"`
	Address   Address     `db:"address"`
	Bank      Bank        `db:"bank"`
	Education []Education `db:"education"`
	Salary    Salary      `db:"salary"`
	Leaves    Leaves      `db:"leaves"`
	Peers     []string    `db:"peers"`
	TeamID    int         `db:"team_id"`
}

// Role for a history of roles.
type Role struct {
	Name  string `db:"name"`
	Start string `db:"start"`
	End   string `db:"end"`
}

//Contact gives the current contact details.
type Contact struct {
	Phone  string `db:"phone"`
	Email  string `db:"email"`
	EPhone string `db:"ephone"`
	EName  string `db:"ename"`
}

// Team details
type Team struct {
	TID     int    `db:"tid"`
	Name    string `db:"name"`
	Manager int    `db:"manager"`
}

// Bank struct contains a users banking account details
type Bank struct {
	PAN     string `db:"pan"`
	Account string `db:"account"`
	IFSC    string `db:"ifsc"`
	Name    string `db:"name"`
}

//Address contains a users Address
type Address struct {
	House    string `db:"house"`
	Street   string `db:"street"`
	District string `db:"district"`
	State    string `db:"state"`
	Zipcode  int    `db:"zipcode"`
	Landmark string `db:"landmark"`
}

//Education struct gives details about a users education qualification
type Education struct {
	ID          int    `db:"id"`
	Institution string `db:"institution"`
	Course      string `db:"course"`
	Yop         int    `db:"yop"`
	Mop         int    `db:"mop"`
}

// Salary structure
type Salary struct {
	Basic int `db:"basic"`
	HRA   int `db:"hra"`
	LTA   int `db:"lta"`
	SPA   int `db:"spa"`
	OR    int `db:"or"`
	PF    int `db:"pf"`
}

type Leaves struct {
	TotalLeaves  int     `db:"totalLeaves"`
	LeaveHistory []Leave `db:"leaveHistory"`
}

type Leave struct {
	Type   string `db:"type"`
	Start  string `db:"start"`
	End    string `db:"end"`
	Status string `db:"status"`
	Reason string `db:"reason"`
}
