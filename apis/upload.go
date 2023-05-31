package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"hechuangfil-admin/models"
)
//批量导入
func UploadAddress(c *gin.Context) {
	var a models.Addr
	file, err := c.FormFile("file")
	if err != nil {

	}
	fmt.Println(file)
	f,_ := file.Open()

	var res models.Response

	if err = a.UploadAddress(f,file.Size);err!=nil{
		res.Msg = err.Error()
		res.Code = 400
		c.JSON(http.StatusOK, res)
	}else{
		res.Msg = "导入成功"
		c.JSON(http.StatusOK, res.ReturnOK())
	}
}
