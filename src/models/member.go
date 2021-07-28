package models

import (
	"log"
	"strings"
	"time"
)

type Member struct {
	account string
	name string
	password string
	locker bool
	locknums int16
	createtime time.Time
}

func (member *Member) UpdatePassword(password string) bool {
	if len(password) <= 6 {
		log.Printf("%s", "密码安全等级过低")
		return false
	}

	if strings.Compare(member.password,password) != 0 {
		member.password = password
	}

	return true
}

func (member Member) GetMember(account string)  {

}