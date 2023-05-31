package apis

import (
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/gin-gonic/gin"
	"hechuangfil-admin/models"
	"hechuangfil-admin/pkg/lotus"
	"net/http"
)
//bls|secp256k1
func WalletNew(c *gin.Context){
	var res models.Response
	t := c.Param("keytype")
	if t==""{
		t = "secp256k1"
	}
	var err error
	res.Data, err = lotus.FullAPI.WalletNew(c, types.KeyType(t))
	if err!=nil{
		c.JSON(http.StatusOK, res.ReturnError(401))
		return
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}