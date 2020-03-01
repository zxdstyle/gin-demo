package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagID uint `json:"tag_id" gorm:"index"`
	Tag Tag	`json:"tag"`
	Title string `json:"title"`
	Desc string	`json:"desc"`
	Content string	`json:"content" gorm:"type:longtext"`
	Status int	`json:"status"`
}


func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}
func GetArticleTotal(maps interface {}) (count int){
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}
func GetArticles(pageNum int, pageSize int, maps interface {}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}
func EditArticle(id int, data interface {}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}
func AddArticle(data map[string]interface {}) bool {
	db.Create(&Article {
		TagID : data["tag_id"].(uint),
		Title : data["title"].(string),
		Desc : data["desc"].(string),
		Content : data["content"].(string),
		Status : data["status"].(int),
	})
	return true
}
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})
	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("createdOn", time.Now().Unix())
	scope.SetColumn("updatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("updatedOn", time.Now().Unix())
	return nil
}

