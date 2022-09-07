package model


type User struct {
	Id int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age int `json:"age"`
}
