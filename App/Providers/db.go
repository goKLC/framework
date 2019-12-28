package Providers

import (
	"fmt"
	"github.com/goKLC/goKLC"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/mssql"
)

func init() {
	var app = goKLC.GetApp()

	dbType := app.Config().Get("DBType", nil)
	dbUrl := app.GetDBURL(dbType.(goKLC.DBType))
	DB, err := gorm.Open(fmt.Sprintf("%v", dbType), dbUrl)

	if err != nil {
		app.Log().Error(err.Error(), nil)
		return
	}

	app.SetDB(DB)
}
