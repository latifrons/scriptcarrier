package rpc

import (
	"github.com/gin-gonic/gin"
	"github.com/latifrons/scriptcarrier/model"
	"github.com/latifrons/scriptcarrier/service"
	"github.com/latifrons/scriptcarrier/tools"
	"net/http"
)

type RpcController struct {
	FolderConfig               tools.FolderConfig
	ReturnDetailedErrorMessage bool
	Mode                       string
	AllowOrigins               []string
	TaskService                *service.TaskService
}

func (rpc *RpcController) ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, GeneralResponse{
		Code: 0,
		Msg:  "",
		Data: data,
	})
}

func (rpc *RpcController) Response(c *gin.Context, status int, code int, msg string, data interface{}) {
	c.JSON(status, GeneralResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func (rpc *RpcController) AddTask(c *gin.Context) {
	var req model.AddTaskRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		rpc.Response(c, http.StatusBadRequest, 1, "bad request", nil)
		return
	}
	err = rpc.TaskService.AddTask(req)
	if err != nil {
		rpc.Response(c, http.StatusConflict, 2, err.Error(), nil)
		return
	}

	rpc.ResponseOK(c, "")
}
