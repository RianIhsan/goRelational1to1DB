package models

type Locker struct {
  ID uint `json:"id" form:"id" gorm:"primaryKey"`
  Code string `json:"code" form:"code" gorm:"unique;not null"`
  UserID int `json:"user_id" form:"user_id" gorm:"not null"`
  User UserRes `json:"user"`
}

type LockerResp struct {
  ID uint `json:"id" form:"id"`
  Code string `json:"code" form:"code"`
  UserID int `json:"-" form:"user_id"`
}

func (LockerResp) TableName() string {
  return "lockers"
}
