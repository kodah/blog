package controller

import (
	"github.com/kodah/blog/service"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBController interface {
	Session() *gorm.DB
}

type dbController struct {
	conn service.DBService
}

func DBHandler() DBController {
	var configService service.ConfigService = service.ConfigurationService("")

	controller := &dbController{
		conn: service.SQLiteDBService(configService.GetDBPath()),
	}

	return controller
}

func (controller dbController) Session() *gorm.DB {
	return controller.conn.Conn().Session(&gorm.Session{
		DryRun:                   false,
		PrepareStmt:              false,
		NewDB:                    false,
		SkipHooks:                false,
		SkipDefaultTransaction:   false,
		DisableNestedTransaction: false,
		AllowGlobalUpdate:        false,
		FullSaveAssociations:     false,
		QueryFields:              false,
		Context:                  nil,
		Logger:                   logger.Default,
		NowFunc:                  nil,
		CreateBatchSize:          0,
	})
}
