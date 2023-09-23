package Keeper

type Keeper struct {
	position string
	salary   float64
	address  string
}

func NewKeeper(position string, salary float64, address string) *Keeper {
	return &Keeper{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (k *Keeper) GetPosition() string {
	return k.position
}

func (k *Keeper) SetPosition(position string) {
	k.position = position
}

func (k *Keeper) GetSalary() float64 {
	return k.salary
}

func (k *Keeper) SetSalary(salary float64) {
	k.salary = salary
}

func (k *Keeper) GetAddress() string {
	return k.address
}

func (k *Keeper) SetAddress(address string) {
	k.address = address
}
