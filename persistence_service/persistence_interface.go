package persistence_service

type PersistenceService interface {
	Open(filePath string, option interface{}) error
	Get(bucketName []byte, key []byte) ([]byte, error)
	Put(bucketName []byte, key []byte, value []byte) error
	Delete(bucketName []byte, key []byte) error
	Close() error
}
