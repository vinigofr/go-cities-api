package cities

type Service interface {
	Create(name, uf string) (city, error)
	GetAll() ([]city, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(name, uf string) (city, error) {}
func (s *service) GetAll() ([]city, error)              {}