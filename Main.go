package main

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"net/http"
)

func main() {
	migrateDB(dba())

	api := iris.New()
	gits_token := api.Party("/gits-token")

	api.Get("/", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK,iris.Map{"status":true,"message":"Hello from the token-service"})
	})
	gits_token.Post("/",createToken)
	gits_token.Put("/",validToken)
	gits_token.Delete("/",deleteToken)
	gits_token.Post("/update",updateToken)

	api.Build()
	fsrv := &http.Server{Handler: api.Router, Addr: ":5000"}
	fsrv.ListenAndServe()
}

//SETUP DB
var dba = func() (*gorm.DB){
	dbs := DBToken{"mysql", "mysql", "21484eb879a418bb","token-service","dokku-mysql-token-service", , nil}
	dbs.getConnection()
	return dbs.dbToken
}

//DB MIGRATE
func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&Token{})
}




