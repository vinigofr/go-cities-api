package cities

type Repository interface {
	Create(name, uf string) (city, error)
	GetAll() ([]city, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repository) Create(name, uf string) (city, error) {
	return city{}, nil
}
func (repository) GetAll() ([]city, error) {
	return []city{}, nil
}