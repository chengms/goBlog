package modules

import "time"

type Article struct {
	Id 			int 	`sql:"not null"`
	Title		string	`sql:"not null"`
	ImgUrl		string
	Content 	string	`sql:"not null"`
	Author 		string 	`sql:"not null"`
	Tag 		string
	Uuid 		int64
	CreateTime 	time.Time	`sql:"not null"`
	Desc 		string
}


