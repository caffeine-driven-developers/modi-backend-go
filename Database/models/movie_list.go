package models

import (
	"github.com/caffeine-driven-developers/modi-backend-go/lib"
	"time"
)

type MovieList struct {

	Movie_List_Id int	 `gorm:"int(11);AUTO_INCREMENT;primary_key"`
	Title		  string `gorm:"varchar(255)"`
	Created_At    time.Time
	Updated_At 	  time.Time
	Movie_Id_List string `gorm:"text"`
}

func (l MovieList) Seriallize() lib.JSON{
	return lib.JSON{
		"movie_list_id": 	l.Movie_List_Id,
		"title": 			l.Title,
		"created_at": 		l.Created_At,
		"updated_at":		l.Updated_At,
		"movie_id_list":	l.Movie_Id_List,
	}
}