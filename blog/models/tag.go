package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"fmt"
)

type Tag struct {
	Model
	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int,pageSize int,maps interface{}) (tags []Tag) {

	db.Debug().Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	fmt.Println(tags[0])
	return
}
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// 新增标签
func ExistTagByName(name string) bool {
	var tag Tag
	db.Debug().Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string,state int,createdBy string) bool {
	db.Create(&Tag{
		Name: name,
		State: state,
		CreatedBy: createdBy,
	})
	return true
}

// 编辑标签
// 检查id是否存在，存在返回true，不存在false
func ExistTagByID(id int) bool {
	var tag Tag
	//  first获取第一条记录，按主键排序
	db.Select("id").Where("id = ?",id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}
// 编辑函数
func EditTag(id int,data interface{}) bool {
	db.Debug().Model(&Tag{}).Where("id = ?",id).Updates(data)
	return true
}

// 删除函数
func DeleteTag(id int) bool {
	db.Where("id = ?",id).Delete(&Tag{})
	return true
}

// 下面是gorm的callback，在创建、更新、查询、删除时将被调用，如果任何回调返回错误，gorm将停止未来操作并回滚所有更改。
// 连接 https://segmentfault.com/a/1190000013297705
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}