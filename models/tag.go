package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	
	Name		 	string 	`json:"name"`
	CreatedBy 		string 	`json:"created_by"`
	ModifiedBy 		string 	`json:"modified_by"`
	State		 	int 	`json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// 查询标签是否存在
func ExistTagByName(name string) bool {
	var tag Tag

	db.Select("id").Where("name = ?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}

	return false
}

// 添加标签
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:       name,
		CreatedBy:  createdBy,
		State:      state,
	})

	return true
}

/* 编写 model callbacks */

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}