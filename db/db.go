package db

import "sync"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Drive struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

func New() {

}

func (d *Drive) Write() error {

}

func (d *Drive) ReadAll() ([]byte, error) {
}

func (d *Drive) Read() ([]byte, error) {
}

func (d *Drive) Delete() error {
}
