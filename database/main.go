package main

import (
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
		log        Logger
		database   string
		collection []string
		mutexes    map[string]*sync.Mutex
		mutex      sync.Mutex
	}
)

func NewDriver(database string, collection []string, log Logger) *Driver {
	return &Driver{
		log:        log,
		database:   database,
		collection: collection,
		mutexes:    make(map[string]*sync.Mutex),
		mutex:      sync.Mutex{},
	}
}

// findOne
// findAll
// findById
// create
// createMany
// update
// updateMany
// delete
// deleteMany

func main() {

}
