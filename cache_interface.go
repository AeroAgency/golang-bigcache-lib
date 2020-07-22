package cache

type CacheInterface interface {
	Get(key string, object interface{}) interface{}
	Set(key string, object interface{}) error
	Delete(key ...string) error
	Clear() error
}
