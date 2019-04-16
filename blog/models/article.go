package models

import (
	"github.com/jinzhu/gorm"
	"time"
	)

type Article struct {
	// gorm:index，用于声明这个字段为索引，如果你使用了自动迁移功能则会有所影响，在不使用则无影响
	// Tag字段，实际是一个嵌套的struct，它利用TagID与Tag模型相互关联，在执行查询的时候，能够达到Article、Tag关联查询的功能
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag Tag `json:"tag"`

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}
// 在article创建后的操作
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	// gettime.Now().Unix() 返回当前的时间戳
	scope.SetColumn("CreateOn",time.Now().Unix())
	return nil
}
// 在article更新后的操作
func (article *Article) BeforUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn",time.Now().Unix())
	return nil
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?",id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

// map 这个稍后看下
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int,pageSize int,maps interface{}) (articles []Article) {
	// Preload就是一个预加载器，它会执行两条SQL，分别是SELECT * FROM blog_articles;和SELECT * FROM blog_tag WHERE id IN (1,2,3,4);，那么在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到Article的Tag中，会特别方便，并且避免了循环查询
	db.Debug().Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
	// 首先执行SELECT * FROM `blog_article`  WHERE (`blog_article`.`state` = '1') LIMIT 2 OFFSET 0;  查询出的tag id用于下面的查询
	// SELECT * FROM `blog_tag`  WHERE (`id` IN ('1','2')) // 这里的id就是上面查询出来的tag id
	// 会返回多个Article结构体 [{{1 0 0} 1 {{1 0 1547971723} edit1 test  0} titletest 描述1 中文测试请问请问  1} {{2 0 0} 2 {{0 0 0}    0} title2 描述2 分公司的分公司的电饭锅电饭锅电饭锅  1}]
}

func GetArticle(id int) (article Article) {
	db.Debug().Where("id = ?",id).First(&article) // 查询传入文章的id值对应的文章，这个值返回给结构体
	db.Debug().Model(&article).Related(&article.Tag) // 根据文章
	return
	// 上面会执行两条sql  SELECT * FROM `blog_article`  WHERE (id = '1') ORDER BY `blog_article`.`id` ASC LIMIT 1
	// SELECT * FROM `blog_tag`  WHERE (`id` = '1') 然后返回结构体
}

func EditArticle(id int,data interface{}) bool {
	db.Model(&Article{}).Where("id = ?",id).Updates(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID: data["tag_id"].(int),
		Title : data["title"].(string),
		Desc : data["desc"].(string),
		Content : data["content"].(string),
		CreatedBy : data["created_by"].(string),
		State : data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})  // 删除操作，没有数据返回传入值类型就行
	return true
}