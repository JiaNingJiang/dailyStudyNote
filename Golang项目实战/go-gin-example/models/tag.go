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
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	tags := make([]Tag, 0)
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, err
}

// 返回tag数据的总数量
func GetTagTotal(maps interface{}) (int, error) {
	var count int
	err := db.Model(&Tag{}).Where(maps).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}

	return count, nil
}
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, err
}

func AddTag(name string, state int, createdBy string) error {
	err := db.Debug().Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}).Error

	return err
}

func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, err
}

func EditTag(id int, data interface{}) error {
	err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error //updates的参数可以是结构体，也可以是map

	return err
}

func DeleteTag(id int) error {
	err := db.Where("id = ?", id).Delete(&Tag{}).Error

	return err
}

func CleanAllTag() error {
	err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{}).Error

	return err
}
