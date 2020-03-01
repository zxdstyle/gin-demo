package v1

import (
	"gin-demo/models"
	"gin-demo/pkg/setting"
	"gin-demo/pkg/util"
	valid "gin-demo/request/v1"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var status int = -1
	if arg := c.Query("status"); arg != "" {
		status = com.StrTo(arg).MustInt()
		maps["status"] = status
	}

	code := util.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c) ,setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  util.GetMsg(code),
		"data": data,
	})


}

//新增文章标签
func AddTag(c *gin.Context) {
	var StoreTagRequest valid.StoreTagRequest

	if err := c.ShouldBind(&StoreTagRequest); err != nil {
		c.JSON(util.INVALID_PARAMS, gin.H{
			"code": 400,
			"msg": err.Error(),
		})
		return
	}

	isPass, err := valid.FormVerify(&StoreTagRequest)
	if !isPass {
		c.JSON(util.INVALID_PARAMS, gin.H{
			"code": 400,
			"msg": err.Error(),
		})
		return
	}

	if ! models.ExistTagByName(StoreTagRequest.Name, 0) {
		models.AddTag(StoreTagRequest.Name, StoreTagRequest.Status)
		c.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" :  util.GetMsg(200),
			"data" : StoreTagRequest,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : 400,
		"msg" :  "该标签已存在",
		"data" : make(map[string]string),
	})
}

//修改文章标签
func EditTag(c *gin.Context) {
	tagID := com.StrTo(c.Param("id")).MustInt()

	var UpdateTagRequest valid.UpdateTagRequest
	if err := c.ShouldBind(&UpdateTagRequest); err != nil {
		c.JSON(util.INVALID_PARAMS, gin.H{
			"code": 400,
			"msg": err.Error(),
		})
		return
	}

	isPass, err := valid.FormVerify(&UpdateTagRequest)
	if !isPass {
		c.JSON(util.INVALID_PARAMS, gin.H{
			"code": 400,
			"msg": err.Error(),
		})
		return
	}

	if ! models.ExistTagByName(UpdateTagRequest.Name, tagID) {
		models.EditTag(tagID, UpdateTagRequest)
		c.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" :  util.GetMsg(200),
			"data" : UpdateTagRequest,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : 400,
		"msg" :  "该标签已存在",
		"data" : make(map[string]string),
	})
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	tagID := com.StrTo(c.Param("id")).MustInt()

	if !models.ExistTagByID(tagID) {
		c.JSON(util.INVALID_PARAMS, gin.H{
			"code": 400,
			"msg": "不存在该标签",
		})
		return
	}

	models.DeleteTag(tagID)

	c.JSON(util.SUCCESS, gin.H{
		"code": 200,
		"msg": "success",
	})
}