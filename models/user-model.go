package models

type User struct {
  ID uint `json:"id" form:"id" gorm:"primaryKey"`
  Name string `json:"name" form:"name" gorm:"not null"`
  Locker LockerResp `json:"locker"` 
}

type UserRes struct {
  ID uint `json:"id" form:"id"`
  Name string `json:"name" form:"name"`
}

func (UserRes) TableName() string {
  return "users"
}
