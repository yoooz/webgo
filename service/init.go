package service

import (
	"errors"
	"fmt"
	"log"
	"webgo/model"

	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := "root:root@tcp([172.17.0.1]:3306)/gin?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}

	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.Book))
	fmt.Println("init data base ok")
}
