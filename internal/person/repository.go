package person

import "github.com/Laanaa/my-app/internal/database"


type Repository interface {
	FindAll() (	[]Person, error)
	FindByID(id int) (*Person, error)
	Create(p Person) (Person, error)
	Update(p Person) (Person, error)
	Delete(id int) error
}

type repo struct{}

func NewRepository() Repository {
	return &repo{}
}

func (r *repo) FindAll() ([]Person, error) {
	var persons []Person
	err := database.DB.Find(&persons).Error
	return persons, err
}

func (r *repo) FindByID(id int) (*Person, error) {
	var person Person
	err := database.DB.First(&person, id).Error
	return &person, err
}

func (r *repo) Create(p Person) (Person, error) {
	err := database.DB.Create(&p).Error
	return p, err
}

func (r *repo) Update(p Person) (Person, error){
	err := database.DB.Save(&p).Error
	return p, err
}

func (r *repo) Delete(id int) error {
	err := database.DB.Delete(&Person{}, id).Error
	return err
}