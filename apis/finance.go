package apis

import (
	"github.com/filecoin-project/go-address"
	"github.com/gin-gonic/gin"
	"hechuangfil-admin/models"
	"hechuangfil-admin/pkg/lotus"
	"hechuangfil-admin/utils"
	"net/http"
)

func FinanceConfigList(c *gin.Context){
	var res models.Response
	var f models.Finance
	fs,err := f.FinanceConfigList()
	if err!=nil{
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	for i:=0;i< len(fs);i++{
		if fs[i].ConfigKey==models.FROM_ADDRESS||fs[i].ConfigKey==models.COLLECTION_ADDRESS{
			addrstr := fs[i].ConfigValue
			addr,err := address.NewFromString(addrstr)
			if err!=nil{
				c.JSON(http.StatusOK, res.ReturnError(400))
				return
			}
			bigInt,err := lotus.FullAPI.WalletBalance(c,addr)
			if err!=nil{
				c.JSON(http.StatusOK, res.ReturnError(400))
				return
			}
			fs[i].Balance = utils.NanoOrAttoToFIL(bigInt.String(),utils.AttoFIL)
		}
	}
	res.Data = fs

	c.JSON(http.StatusOK, res.ReturnOK())
}

func UpdateConfigById(c *gin.Context) {
	var res models.Response
	var f models.Finance
	param := make(map[string]string)
	param["id"] = c.Request.FormValue("id")
	param["title"] = c.Request.FormValue("title")
	param["value"] = c.Request.FormValue("value")
	fs,err := f.UpdateConfigById(param)
	if err!=nil{
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Data = fs
	c.JSON(http.StatusOK,res.ReturnOK())
}