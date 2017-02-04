package main

import (
	"time"
	"github.com/jinzhu/gorm"
)
//MODEL DATABASE TOKEN
type Token struct {
	gorm.Model
	Apps string //APPS CODE
	UserID string //USER ID IF SET
	Token string //TOKEN GENERATED
}

//INSERT TOKEN
func (mt *Token) insertToken(dba *gorm.DB) *Token{
	defer dba.Close()
	mt.Token = GetMD5Hash(mt.Apps+":"+time.Now().String())
	dba.NewRecord(mt)
	dba.Create(&mt)
	return mt
}
//UPDATE TOKEN
func (mt *Token) updateToken(dba *gorm.DB,data Token) (*Token,bool,string)  {
	defer dba.Close()
	var mba = Token{}
	if err := dba.Where(&data).First(&mba).Error; err != nil {
		return &mba,false,err.Error()
	}else{
		mba.UserID = mt.UserID
		dba.Save(&mba)
		return &mba,true,"success update"
	}
}
//DELETE TOKEN
func (mt *Token) deleteToken(dba *gorm.DB) (token *Token,status bool,message string)   {
	defer dba.Close()
	var mba = Token{}
	if err := dba.Where(&mt).First(&mba).Error; err != nil {
		token = &mba
		status = false
		message = err.Error()
	}else{
		dba.Delete(&mt)
		token = &mba
		status = true
		message = "token deleted"
	}
	return
}
//CHECK TOKEN
func (mt *Token) checkToken(dba *gorm.DB) (token *Token,status bool,message string) {
	defer dba.Close()
	var mba = Token{}
	if err := dba.Where(&mt).Find(&mba).Error; err != nil {
		token = &mba
		status = false
		message = "token not valid"
	}else {
		token = &mba
		status = true
		message = "token valid"
	}
	return
}