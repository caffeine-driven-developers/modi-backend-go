package models

import "time"

type MovieList struct {

	Movie_List_Id int	 `gorm:"AUTO_INCREMENT;primary_key"`
	Title		  string `gorm:"varchar(255)"`
	Created_At    time.Time
	Updated_At 	  time.Time
	Movie_Id_List string `gorm:"text"`
}
