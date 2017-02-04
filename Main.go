package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func main() {
	fmt.Print("This Service For Token")
	migrateDB(dba())

	token := Token{}
	token.Token = "13f939b6e33316db7e2306a058c22bb0"
	status,message := token.checkToken(dba())
	println(status)
	println(message)
}

//SETUP DB
var dba = func() (*gorm.DB){
	dbs := DBToken{"mysql", "root", "", "service_token", nil}
	dbs.getConnection()
	return dbs.dbToken
}

//DB MIGRATE
func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&Token{})
}




