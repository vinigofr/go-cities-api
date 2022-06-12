package cities

type Repository interface {
	Create(name, uf string) (City, error)
	GetAll() ([]City, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repository) Create(name, uf string) (City, error) {
	return City{}, nil
}
func (repository) GetAll() ([]City, error) {
	return []City{}, nil
}