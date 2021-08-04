package models

import (
	"log"
	"strings"
	"time"
)

type User interface {
	SetUser(member Member)
	GetUser() Member
}

type Member struct {
	Id interface{}
	Account string
	Name string
	Password string
	Locker bool
	Locknums int16
	Createtime time.Time
}

func (member *Member) UpdatePassword(password string) bool {
	if len(password) <= 6 {
		log.Printf("%s", "密码安全等级过低")
		return false
	}

	if strings.Compare(member.Password,password) != 0 {
		member.Password = password
	}

	return true
}