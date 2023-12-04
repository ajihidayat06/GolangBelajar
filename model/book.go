package model

type Book struct {
	Id     int64  `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title  string `json:"title" validate:"required"`
	Author int64  `json:"author"`
	Desc   string `json:"desc" validate:"required"`
}

type BookResponse struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}
