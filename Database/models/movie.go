package models

import "time"

type Movie struct {
	Movie_Id     string 	`gorm:"char(16);primary_key"`
	Title        string 	`gorm:"varchar(255)"`
	Year         string 	`gorm:"varchar(255)"`
	Running_Time string 	`gorm:"varchar(255)"`
	Genre        string 	`gorm:"varchar(255)"`
	Director     string 	`gorm:"varchar(255)"`
	Imdb_Rating  string 	`gorm:"varchar(255)"`
	Actors       string 	`gorm:"varchar(255)"`
	Created_At   time.Time

}
