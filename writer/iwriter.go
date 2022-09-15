package writer

type IWriter interface {
	Write(filepath string) error
}
