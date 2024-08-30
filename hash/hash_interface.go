package hash

type Hash interface {
	Name() string
	Hash(data []byte) uint64
}
