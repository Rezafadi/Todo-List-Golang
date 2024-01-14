package models

type List struct {
	Id          int       `json:"id" gorm:"column:id;primary_key;auto_increment"`
	Title       string    `json:"title" gorm:"column:title;not null"`
	Description string    `json:"description" gorm:"column:description;type:text;size:1000;not null"`
	File        string    `json:"file"`
	SubLists    []SubList `json:"sub_lists" gorm:"foreignKey:ListID;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;preload:SubLists"`
}

func (List) TableName() string {
	return "lists"
}