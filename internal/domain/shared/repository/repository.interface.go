package shared

type RepositoryInterface[T any] interface {
	Create(entity *T) error
	Get(page int) ([]*T, error)
	GetById(id string) (*T, error)
	Delete(entity *T) error
}
