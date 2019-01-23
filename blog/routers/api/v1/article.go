package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/learn-go/blog/pkg/e"
	"github.com/learn-go/blog/models"
	"log"
	"net/http"
	"github.com/learn-go/blog/pkg/util"
	"github.com/learn-go/blog/pkg/setting"
)

//  获取单个文章
func GetArticle(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		}else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}else {
		for _,err := range valid.Errors {
			log.Println(err.Key,err.Message)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

// 获取多个文章
func GetArticles(c *gin.Context)  {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state,0,1,"state").Message("状态只允许0或1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetArticles(util.GetPage(c),setting.PageSize,maps)
		data["total"] = models.GetTagTotal(maps)
	} else {
		for _,err := range valid.Errors{
			log.Println(err.Key,err.Message)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

// 请求参数
// {"tag_id": 1,
//"title": "test1",
//"desc": "test-desc",
//"content": "test-content",
//"created_by": "test-created",
//"state": 1
//}
type addArticlesBody struct {
	State int `json:"state"`
	TagId int `json:"tag_id"`
	Title string `json:"title"`
	CreatedBy string `json:"created_by"`
	Content string `json:"content"`
	Desc string `json:"desc"`

}
// 新增文章
func AddArticle(c *gin.Context)  {
	//tagId := com.StrTo(c.Query("tag_id")).MustInt()
	//title := c.Query("title")
	//desc := c.Query("desc")
	//content := c.Query("content")
	//createdBy := c.Query("created_by")
	//state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	var addArticlesBody addArticlesBody
	err := c.BindJSON(&addArticlesBody)
	if err != nil {
		log.Println(err)
		// 这个位置有警告回头看下22322222222 [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 200
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}

	//valid := validation.Validation{}
	//valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	//valid.Required(title, "title").Message("标题不能为空")
	//valid.Required(desc, "desc").Message("简述不能为空")
	//valid.Required(content, "content").Message("内容不能为空")
	//valid.Required(createdBy, "created_by").Message("创建人不能为空")
	//valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	valid := validation.Validation{}
	valid.Min(addArticlesBody.TagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(addArticlesBody.Title, "title").Message("标题不能为空")
	valid.Required(addArticlesBody.Desc, "desc").Message("简述不能为空")
	valid.Required(addArticlesBody.Content, "content").Message("内容不能为空")
	valid.Required(addArticlesBody.CreatedBy, "created_by").Message("创建人不能为空")
	valid.Range(addArticlesBody.State, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistTagByID(addArticlesBody.TagId) {
			data := make(map[string]interface {})
			data["tag_id"] = addArticlesBody.TagId
			data["title"] = addArticlesBody.Title
			data["desc"] = addArticlesBody.Desc
			data["content"] = addArticlesBody.Content
			data["created_by"] = addArticlesBody.CreatedBy
			data["state"] = addArticlesBody.State

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key,err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}

// 修改文章
func EditArticle(c *gin.Context)  {
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagByID(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}
				data["modified_by"] = modifiedBy

				models.EditArticle(id,data)
				code = e.SUCCESS
			}else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		}else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}else {
		for _, err := range valid.Errors {
			log.Println(err.Key,err.Message)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}