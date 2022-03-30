package repo

type Repository interface {
	Create(object interface{}) error
	Read(object interface{}) error
	Update(object interface{}) error
	Delete(object interface{}) error
}
