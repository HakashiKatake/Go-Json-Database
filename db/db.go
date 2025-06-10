package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (d *Driver) Write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing collection - no place to save record!")
	}

	if resource == "" {
		return fmt.Errorf("Missing resource - no place to save record!")
	}

	mutex := d.getOrCreateMutex()
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	fnlPath := filepath.Join(dir, resource+".json")
	tmpPath := fnlPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", dir, err)
	}

	b, err := json.MarshalIndent(v, "", " \t")
	if err != nil {
		return err
	}

	b = append(b, byte('\n'))
	if err := ioutil.WriteFile(tmpPath, b, 0644); err != nil {
		return fmt.Errorf("Failed to write to temporary file %s: %w", tmpPath, err)
	}

}

func (d *Driver) ReadAll() ([]byte, error) {
}

func (d *Driver) Read() ([]byte, error) {
}

func (d *Driver) Delete() error {
}

func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {

}

func stat(path string) (fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return
}
