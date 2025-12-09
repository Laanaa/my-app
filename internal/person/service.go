package person

type Service interface {
	GetAll() ([]Person, error)
	GetByID(id int) (*Person, error)
	Create(p Person) (Person, error)
	Update(id int, p Person) (Person, error)
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]Person, error) {
	return s.repo.FindAll()
}

func (s *service) GetByID(id int) (*Person, error) {
	return s.repo.FindByID(id)
}

func (s *service) Create(p Person) (Person, error) {
	return s.repo.Create(p)
}

func (s *service) Update(id int, p Person) (Person, error) {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return p, err
	}

	existing.Name = p.Name
	existing.Age = p.Age

	return s.repo.Update(*existing)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}