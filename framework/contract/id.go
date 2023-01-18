package contract

const IDKey = "beide:id"

type IDService interface {
	NewID() string
}
