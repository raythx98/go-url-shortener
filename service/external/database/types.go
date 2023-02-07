package database

type Config struct {
	Host    string
	Timeout int
}

type Client interface {
	Close() error
}
