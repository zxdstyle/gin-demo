package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name string	`json:"name"`
	Status int	`json:"status"`
}

func GetTags(page int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(page).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// 判断是否已存在某标签 tagID为忽略检查的标签ID
func ExistTagByName(name string, tagId int) bool {
	var tag Tag
	query := db.Select("id").Where("name = ?", name)

	if tagId > 0 {
		query.Where("id <> ?", tagId)
	}

	query.First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, status int) bool {
	db.Create(&Tag{
		Name: name,
		Status: status,
	})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("createdOn", time.Now().Unix())
	scope.SetColumn("updatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("updatedOn", time.Now().Unix())
	return nil
}

