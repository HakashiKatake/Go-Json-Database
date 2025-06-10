package db

import (
	"os"
	"path/filepath"
	"sync"
)

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Database directory already exists, using existing directory: %s", dir)
		return &driver, nil
	}

	opts.Logger.Debug("Creating database directory: %s", dir)
	return &driver, os.MkdirAll(dir, 0755)

}

func (d *Driver) Write() error {

}

func (d *Driver) ReadAll() ([]byte, error) {
}

func (d *Driver) Read() ([]byte, error) {
}

func (d *Driver) Delete() error {
}

func (d *Driver) getOrCreateMutex() {

}
