package reader

type IReader interface {
	Read() ([]byte, error)
}
