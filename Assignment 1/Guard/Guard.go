package Guard

type Guard struct {
	position string
	salary   float64
	address  string
}

func NewGuard(position string, salary float64, address string) *Guard {
	return &Guard{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (g *Guard) GetPosition() string {
	return g.position
}

func (g *Guard) SetPosition(position string) {
	g.position = position
}

func (g *Guard) GetSalary() float64 {
	return g.salary
}

func (g *Guard) SetSalary(salary float64) {
	g.salary = salary
}

func (g *Guard) GetAddress() string {
	return g.address
}

func (g *Guard) SetAddress(address string) {
	g.address = address
}
