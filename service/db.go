package service

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbStore DBService
)

type DBService interface {
	Error() error
	Conn() *gorm.DB
}

type db struct {
	conn *gorm.DB

	error error

	mutex sync.RWMutex
}

func SQLiteDBService(path string) DBService {
	if dbStore == nil {
		conn, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

		service := new(db)
		service.conn = conn
		service.error = err
		// make the service persistent
		dbStore = service

		return service
	}

	return dbStore
}

// Used for detecting errors after loading
func (d *db) Error() error {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	return d.error
}

func (d *db) Conn() *gorm.DB {
	return d.conn
}
