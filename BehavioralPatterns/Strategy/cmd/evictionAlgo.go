package cmd

type EvictionAlgo interface {
	evict(c *Cache)
}
