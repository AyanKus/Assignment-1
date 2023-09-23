package Runner

type Runner struct {
	position string
	salary   float64
	address  string
}

func NewRunner(position string, salary float64, address string) *Runner {
	return &Runner{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (r *Runner) GetPosition() string {
	return r.position
}

func (r *Runner) SetPosition(position string) {
	r.position = position
}

func (r *Runner) GetSalary() float64 {
	return r.salary
}

func (r *Runner) SetSalary(salary float64) {
	r.salary = salary
}

func (r *Runner) GetAddress() string {
	return r.address
}

func (r *Runner) SetAddress(address string) {
	r.address = address
}
