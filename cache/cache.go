package cache

type Cache interface {
	SetKey(Model) error
	GetKey(Model) error
}
