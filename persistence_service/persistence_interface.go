package persistence_service

type PersistenceService interface {
	Open(filePath string, option interface{}) error
	Get(key []byte) ([]byte, error)
	Put(key []byte, value []byte) error
	Delete(key []byte) error
	Close() error
}
