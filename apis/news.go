package apis

import (
	"github.com/gin-gonic/gin"
	"hechuangfil-admin/models"
	"hechuangfil-admin/result"
	"hechuangfil-admin/utils"
	"net/http"
)

func GetAllNews(c *gin.Context) {
	pageSize := c.Request.FormValue("pageSize")
	if pageSize == "" {
		c.JSON(http.StatusOK, result.Failstr("pageSize 不能为空"))
		c.Abort()
		return
	}
	pageIndex := c.Request.FormValue("pageIndex")
	if pageIndex == "" {
		c.JSON(http.StatusOK, result.Failstr("pageIndex 不能为空"))
		return
	}

	news := models.NewNews()
	param := make(map[string]string)
	param["pageIndex"] = pageIndex
	param["pageSize"] = pageSize
	ret, err := news.GetAll(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	res := utils.NewPageData(param, ret)
	c.JSON(http.StatusOK, res)
}

func AddNews(c *gin.Context) {
	title := c.Request.FormValue("title")
	if title == "" {
		c.JSON(http.StatusOK, result.Failstr("title 不能为空"))
		c.Abort()
		return
	}
	content := c.Request.FormValue("content")
	if content == "" {
		c.JSON(http.StatusOK, result.Failstr("content 不能为空"))
		c.Abort()
		return
	}
	status := c.Request.FormValue("status")
	if status == "" {
		c.JSON(http.StatusOK, result.Failstr("status 不能为空"))
		c.Abort()
		return
	}

	news := models.NewNews()
	err := news.AddNews(title, content, status)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(nil))
}

func UpdateNews(c *gin.Context) {
	id := c.Request.FormValue("id")
	if id == "" {
		c.JSON(http.StatusOK, result.Failstr("id 不能为空"))
		c.Abort()
		return
	}
	title := c.Request.FormValue("title")
	if title == "" {
		c.JSON(http.StatusOK, result.Failstr("title 不能为空"))
		c.Abort()
		return
	}
	content := c.Request.FormValue("content")
	if content == "" {
		c.JSON(http.StatusOK, result.Failstr("content 不能为空"))
		c.Abort()
		return
	}
	status := c.Request.FormValue("status")
	if status == "" {
		c.JSON(http.StatusOK, result.Failstr("status 不能为空"))
		c.Abort()
		return
	}

	news := models.NewNews()
	err := news.Update(id, title, content, status)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(nil))
}

func DeleteNews(c *gin.Context) {
	id := c.Request.FormValue("id")
	if id == "" {
		c.JSON(http.StatusOK, result.Failstr("id 不能为空"))
		c.Abort()
		return
	}
	news := models.NewNews()
	if err := news.Del(id); err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(nil))
}

func UpDateNewStatus(c *gin.Context) {
	id := c.Request.FormValue("id")
	if id == "" {
		c.JSON(http.StatusOK, result.Failstr("id 不能为空"))
		c.Abort()
		return
	}

	status := c.Request.FormValue("status")
	if status == "" {
		c.JSON(http.StatusOK, result.Failstr("status 不能为空"))
		c.Abort()
		return
	}

	news := models.NewNews()
	if err := news.UpdateStatus(id, status); err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Ok(nil))
}
