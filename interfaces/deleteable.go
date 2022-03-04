package interfaces

type Deleteable interface {
	Delete() (bool, error)
}
