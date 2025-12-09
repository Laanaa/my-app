package person

type Service interface {
	GetAll() []Person
	GetByID(id int) *Person
	Create(p Person) Person
	Update(id int, p Person) *Person
	Delete(id int) bool
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() []Person {
	return s.repo.FindAll()
}

func (s *service) GetByID(id int) *Person {
	return s.repo.FindByID(id)
}

func (s *service) Create(p Person) Person {
	return s.repo.Create(p)
}

func (s *service) Update(id int, p Person) *Person {
	return s.repo.Update(id, p)
}

func (s *service) Delete(id int) bool {
	return s.repo.Delete(id)
}