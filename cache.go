package abstraction

import "time"

type Cache interface {
	InitCache()
	Store(key string, data interface{}, duration time.Duration) error
	Exists(key string) bool
	Get(key string) ([]byte, error)
	Delete(key string) error
}
