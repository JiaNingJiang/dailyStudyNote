package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct { // 无论是在查询还是插入时，Tag类的对象都会直接在blog_tag表中完成映射(前缀blog_是自行指定的)
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error { //数据项被创建之前被调用
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error { //数据项被更新之前被调用
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

// 返回从pageNum开始的pageSize条tag数据
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

// 返回tag数据的总数量
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func AddTag(name string, state int, createdBy string) bool {
	db.Debug().Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data) //updates的参数可以是结构体，也可以是map

	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

	return true
}
