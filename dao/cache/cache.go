package cache

type Cache struct {
	CacheType string
	Redis     []cache.RedisConf
}

var cache Cache

func (cache *Cache) get() {

}

func (cache *Cache) set() {

}
