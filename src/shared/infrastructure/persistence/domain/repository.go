package persistence_domain

type Repository[T any] interface {
	FindBy(id int) T
	All() []T
}
