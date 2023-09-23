package Administrator

type Administrator struct {
	position string
	salary   float64
	address  string
}

func NewAdministrator(position string, salary float64, address string) *Administrator {
	return &Administrator{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (a *Administrator) GetPosition() string {
	return a.position
}

func (a *Administrator) SetPosition(position string) {
	a.position = position
}

func (a *Administrator) GetSalary() float64 {
	return a.salary
}

func (a *Administrator) SetSalary(salary float64) {
	a.salary = salary
}

func (a *Administrator) GetAddress() string {
	return a.address
}

func (a *Administrator) SetAddress(address string) {
	a.address = address
}
