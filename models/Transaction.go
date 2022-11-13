package models

import "time"

type Transaction struct {
	ID           int          `json:"id" gorm:"primary_key:auto_increment" `
	UserFundID   int          `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
	UserFund     UserResponse `json:"user_fund" `
	UserDonateID int          `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"  `
	UserDonate   UserResponse `json:"user_donate"  `
	FundID       int          `json:"fund_id"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
	Fund         Fund         `json:"fund" `
	Donate       int          `json:"donate"`
	Status       string       `json:"status"`
	CreatedAt    time.Time    `json:"create_at"`
	UpdatedAt    time.Time    `json:"-"`
}
