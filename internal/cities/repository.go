package cities

type Repository interface {
	Create(name, uf string) (City, error)
	GetAll() ([]City, error)
}

type repository struct{}

var Database []City
var globalID = 1

func NewRepository() Repository {
	return &repository{}
}

func (repository) Create(name, uf string) (City, error) {
	newCity := City{Id: globalID, Name: name, UF: uf}
	Database = append(Database, newCity)
	globalID++

	return newCity, nil
}
func (repository) GetAll() ([]City, error) {
	return Database, nil
}