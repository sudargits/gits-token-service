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
func (mt *Token) insertToken(dba *gorm.DB) bool{
	defer dba.Close()
	mt.Token = GetMD5Hash(mt.Apps+":"+time.Now().String())
	isCreate := dba.NewRecord(mt)
	dba.Create(&mt)
	return isCreate
}
//UPDATE TOKEN
func (mt *Token) updateToken(dba *gorm.DB,data Token) *Token  {
	dba.First(&mt)
	defer dba.Close()
	mt.UserID = data.UserID
	mt.Apps = data.Apps

	dba.Save(&mt)
	return mt
}
//DELETE TOKEN
func (mt *Token) deleteToken(dba *gorm.DB) bool  {
	defer dba.Close()
	dba.First(&mt)
	dba.Delete(&mt)
	return true
}
//CHECK TOKEN
func (mt *Token) checkToken(dba *gorm.DB) (status bool,message string) {
	defer dba.Close()
	var mba = Token{}
	if err := dba.Where(&mt).Find(&mba).Error; err != nil {
		status = false
		message = "token not valid"
	}else {
		status = true
		message = "token valid"
	}
	return
}