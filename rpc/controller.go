package rpc

import (
	"github.com/gin-gonic/gin"
	"github.com/latifrons/scriptcarrier/tools"
	"net/http"
)

type RpcController struct {
	FolderConfig               tools.FolderConfig
	ReturnDetailedErrorMessage bool
	Mode                       string
	AllowOrigins               []string
}

func (rpc *RpcController) ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, GeneralResponse{
		Code: 0,
		Msg:  "",
		Data: data,
	})
}
