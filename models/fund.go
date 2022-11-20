package models

import "time"

type Fund struct {
	ID        int          `json:"id" gorm:"primary_key:auto_increment"`
	Name      string       `json:"name" gorm:"type:varchar(255)"`
	Image     string       `json:"image" gorm:"type:varchar(255)"`
	Desc      string       `json:"desc" gorm:"type:text"`
	Goals     int          `json:"goals" gorm:"type:int"`
	Donated   int          `json:"donated" gorm:"type:int"`
	UserID    int          `json:"-" `
	User      UserResponse `json:"user"`
	Date      time.Time    `json:"date" gorm:"default:now()" `
	CreatedAt time.Time    `json:"create_at"`
	Update_at time.Time    `json:"-"`
}

type FundResponse struct {
	ID      int          `json:"id"`
	Name    string       `json:"name"`
	Image   string       `json:"image"`
	Desc    string       `json:"desc"`
	Goals   int          `json:"goals"`
	Donated int          `json:"donated"`
	UserID  int          `json:"-"`
	User    UserResponse `json:"user"`
}

func (FundResponse) TableName() string {
	return "funds"
}
