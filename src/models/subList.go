package models

type SubList struct {
	Id          int    `json:"id" gorm:"column:id;primary_key;auto_increment"`
	Title       string `json:"title" gorm:"column:title;not null"`
	Description string `json:"description" gorm:"column:description;type:text;size:1000;not null"`
	File        string `json:"file"`
	ListID      int    `json:"list_id" gorm:"column:list_id;not null;foreignKey:Id"`
}

func (SubList) TableName() string {
	return "sub_lists"
}