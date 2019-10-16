package cache

// Cache is an interface, which can be used for Repository caching
type Cache interface {
	Write(string)
	ReadAll() []string
}
