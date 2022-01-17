package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/vino42/go-gin-example/models"
	"github.com/vino42/go-gin-example/pkg/e"
	"github.com/vino42/go-gin-example/pkg/setting"
	"github.com/vino42/go-gin-example/pkg/util"
	"net/http"
)

func GetTag(c *gin.Context) {
	var state int = 1
	maps := make(map[string]interface{})

	id := c.Query("id")

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	maps["id"] = id

	code := e.SUCCESS

	data := models.GetTag(maps, 0, nil)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func GetTags(c *gin.Context) {
	name := c.Query("name")

	var state int = 1
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS
	data["records"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 接收postbody json请求 https://www.cnblogs.com/devhg/p/12776270.html
func AddTag(c *gin.Context) {
	var t models.Tag
	if c.Bind(&t) == nil {
		name := t.Name
		createdBy := t.CreatedBy
		state := t.State

		valid := validation.Validation{}
		valid.Required(name, "name").Message("名称不能为空")
		valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
		valid.Required(createdBy, "created_by").Message("创建人不能为空")
		valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

		code := e.INVALID_PARAMS
		if !valid.HasErrors() {
			if !models.ExistTagByName(name) {
				code = e.SUCCESS
				models.AddTag(name, state, createdBy)
			} else {
				code = e.ERROR_EXIST_TAG
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]string),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  "请求参数错误",
		"data": make(map[string]string),
	})
}

func EditTag(c *gin.Context) {
	var t models.Tag
	if c.Bind(&t) == nil {
		name := t.Name
		modifiedBy := t.ModifiedBy
		state := t.State
		id := t.ID

		valid := validation.Validation{}
		valid.Required(name, "name").Message("名称不能为空")
		valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
		valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
		valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

		code := e.INVALID_PARAMS
		if !valid.HasErrors() {
			if models.ExistTagByID(id) {
				data := make(map[string]interface{})
				data["modified_by"] = modifiedBy
				if name != "" {
					data["name"] = name
				}
				if state != -1 {
					data["state"] = state
				}

				code = e.SUCCESS
				models.EditTag(id, data)
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]string),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  "请求参数错误",
		"data": make(map[string]string),
	})
}

func DeleteTag(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := com.StrTo(id).Int()
	models.DeleteTag(idInt)
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": make(map[string]string),
	})
}
