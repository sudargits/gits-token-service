package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type DBToken struct {
	type_db string // TIPE DATABASE
	username string // USERNAME DATA ASE
	password string // PASSWORD
	db_name string // DATABASE NAME
	db_host string
	dbToken *gorm.DB // PELENGKAP
}
//CONNECTION DB
func (dbToken *DBToken) getConnection() *gorm.DB {
	var s = dbToken.username+":"+dbToken.password+"@"+dbToken.db_host+"/"+dbToken.db_name+"?charset=utf8&parseTime=True&loc=Local"
	fmt.Print(s)

	dba, err := gorm.Open(dbToken.type_db, dbToken.username+":"+dbToken.password+"@"+dbToken.db_host+"/"+dbToken.db_name+"?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		dbToken.dbToken = dba
	}else{
		panic(err.Error())
	}
	return dba
}
//DB DESTROY
func (dbToken *DBToken) destroy(connection *gorm.DB){
	connection.Close()
}
