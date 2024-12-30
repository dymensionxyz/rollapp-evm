package repository

type Repository interface {
	Get(key []byte) ([]byte, error)
	Save(key, value []byte) error
	Close() error
}
