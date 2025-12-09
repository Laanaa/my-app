package person

var persons = []Person{}
var lastID = 0

type Repository interface {
	FindAll() []Person
	FindByID(id int) *Person
	Create(p Person) Person
	Update(id int, p Person) *Person
	Delete(id int) bool
}

type repo struct{}

func NewRepository() Repository {
	return &repo{}
}

func (r *repo) FindAll() []Person {
	return persons[:]
}

func (r *repo) FindByID(id int) *Person {
	for i := range persons {
		if persons[i].ID == id {
			return &persons[i]
		}
	}
	return nil
}

func (r *repo) Create(p Person) Person {
	lastID++
	p.ID = lastID
	persons = append(persons, p)
	return p
}

func (r *repo) Update(id int, p Person) *Person {
	for i := range persons {
		if persons[i].ID == id {
			persons[i].Name = p.Name
			persons[i].Age = p.Age
			return &persons[i]
		}
	}
	return nil
}

func (r *repo) Delete(id int) bool {
	for i := range persons {
		if persons[i].ID == id {
			persons = append(persons[:i], persons[i+1:]...)
			return true
		}
	}
	return false
}