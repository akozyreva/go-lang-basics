package magazine

type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	sex         string  // only fields from capital letters are being exported
	HomeAddress Address // we can use another structs as type
	ContactInfo
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
	ContactInfo // I set it anonymously, because otherwise it's too ambiguous
}

type Address struct {
	Street   string
	City     string
	State    string
	Postcode int
}

type ContactInfo struct {
	PhoneNumber string
	Email       string
}
