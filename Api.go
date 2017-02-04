package main

import (
	"github.com/kataras/iris"
)

type TokenResponse struct {
	Token string `json:"token"`
	Apps string `json:"apps"`
	UserID string `json:"user_id"`
}
type MessageResponse struct {
	Devel string `json:"devel"`
	Prod string  `json:"prod"`
}
type TokenREST struct {
	Status bool `json:"status"`
	Message MessageResponse `json:"message"`
	Content TokenResponse `json:"content"`
}
func createToken(ctx *iris.Context) {
	token := Token{}
	token.Apps = ctx.FormValue("apps")
	token.insertToken(dba())

	var response =  TokenREST{
		Status:true,
		Message:MessageResponse{Devel:"Success create token",Prod:"Success create token"},
		Content:TokenResponse{Apps:token.Apps,Token:token.Token},
	}
	ctx.JSON(iris.StatusOK,response)
}

func updateToken(ctx *iris.Context) {
	token := Token{}
	token.Apps = ctx.FormValue("apps")
	token.Token = ctx.FormValue("token")
	token.UserID = ctx.FormValue("userid")

	find := Token{}
	find.Apps = ctx.FormValue("apps")
	find.Token = ctx.FormValue("token")

	data,status,message := token.updateToken(dba(),find)
	var response  = TokenREST{
		Status:status,
		Message:MessageResponse{Devel:message,Prod:message},
		Content:TokenResponse{Apps:data.Apps,Token:data.Token,UserID:data.UserID},
	}
	ctx.JSON(iris.StatusOK,response)
}

func validToken(ctx *iris.Context) {
	token := Token{}
	token.Token = ctx.FormValue("token")
	token.Apps = ctx.FormValue("apps")
	data,status,message  := token.checkToken(dba())

	var messageProd = ""
	if status {
		messageProd = "Token is valid"
	}else {
		messageProd = "Token not validr"
	}

	var response  = TokenREST{
		Status:status,
		Message:MessageResponse{Devel:message,Prod:messageProd},
		Content:TokenResponse{Apps:data.Apps,Token:data.Token,UserID:data.UserID},
	}
	ctx.JSON(iris.StatusOK,response)
}

func deleteToken(ctx *iris.Context) {
	token := Token{}
	token.Token = ctx.FormValue("token")
	token.Apps = ctx.FormValue("apps")
	data,status,messageDevel := token.deleteToken(dba())

	var messageProd = ""
	if status {
		messageProd = "Token expired now"
	}else {
		messageProd = "System error, please try again"
	}
	var response =  TokenREST{
		Status:status,
		Message:MessageResponse{Devel:messageDevel,Prod:messageProd},
		Content:TokenResponse{Apps:data.Apps,Token:data.Token,UserID:data.UserID},
	}

	ctx.JSON(iris.StatusOK,response)
}