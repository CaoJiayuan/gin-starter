package models

type Article struct {
	ID      int64  `gorm:"column:id;primaryKey;not null;type:int;autoIncrement" json:"id" form:"id"`
	Title   string `gorm:"column:title;type:varchar(255)" json:"title" bind:"required" form:"title"`
	Content string `gorm:"column:content;type:text" json:"content" bind:"required" form:"content"`
}
