package types

// give me a interface of db client

type Service interface {
	interface {
		Init() error
		Close() (func(), error)
	}
}
