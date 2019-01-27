package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/learn-go/blog/pkg/e"
	"github.com/learn-go/blog/models"
		"github.com/learn-go/blog/pkg/setting"
	"github.com/learn-go/blog/pkg/util"
	"net/http"
	"github.com/astaxie/beego/validation"
	"fmt"
)

// 获取多个文章标签
func GetTags(c *gin.Context) {
	//c.Query可用于获取?name=test&state=1这类URL参数，而c.DefaultQuery则支持设置一个默认值
	name := c.Query("name")

	// 定义map
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		// 使用com包里带的函数，强制转成int
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c),setting.PageSize, maps)
	fmt.Println("data的值为：",data)
	data["total"] = models.GetTagTotal(maps)

	// 返回json信息，还有data数据
	c.JSON(http.StatusOK,gin.H{
		"code": code,
		// 获取状态码对应的信息
		"msg": e.GetMsg(code),
		"data": data,
	})
}

// 新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	// 这个是默认state不存在的话，就给一个state=0的值
	state := com.StrTo(c.DefaultQuery("state","0")).MustInt()
	createdBy := c.Query("created_by")

	// 参数验证，用的beego提供的
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name,100,"name").Message("名称最长为100字符")
	valid.Required(createdBy,"created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy,100,"created_by").Message("创建人最长为100字符")
	valid.Range(state,0,1,"state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	// 如果上面的验证不错的话继续
	if ! valid.HasErrors() {
		// 判断数据库中是否有同名的name，没有的话code变成200
		if ! models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name,state,createdBy)
		}else {
			code = e.ERROR_EXIST_TAG
		}
	}

	// 这个data目前返回为空，不知道作用
	c.JSON(http.StatusOK,gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	// 检测state的值，预定义为-1，不定义的话值默认为0会有问题
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state,0	,1,"state").Message("状态只允许0或1")
	}

	valid.Required(id,"id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	// 没有错误的话把状态码改为200
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			// 定义map，存入url传来的数据
			data := make(map[string]interface{})
			// 增加修改者名称
			// data里的字段名，要和数据库的字段名对上
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			fmt.Println("更新的数据为：",data)
			models.EditTag(id,data)
		} else {
			// 不存在返回相应错误码
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 删除文章标签
func DleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		}else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})

}