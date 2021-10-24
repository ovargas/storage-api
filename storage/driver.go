package storage

import (
	"github.com/ovargas/storage-api/storage/driver"
	"sync"
)


var (
	driversMu sync.RWMutex
	drivers   = make(map[string]driver.Driver)
)

func Register(driverName string, driver driver.Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("storage: Register driver is nil")
	}
	if _, dup := drivers[driverName]; dup {
		panic("storage: Register called twice for provider " + driverName)
	}
	drivers[driverName] = driver
}
