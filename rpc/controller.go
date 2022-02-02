package rpc

import (
	"github.com/gin-gonic/gin"
	"github.com/latifrons/scriptcarrier/db"
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

func (rpc *RpcController) ResponsePaging(c *gin.Context, pagingResult model.PagingResult, data interface{}, list interface{}) {
	c.JSON(http.StatusOK, PagingResponse{
		GeneralResponse: GeneralResponse{
			Code: 0,
			Msg:  "",
			Data: data,
		},
		List:  list,
		Size:  pagingResult.Limit,
		Total: pagingResult.Total,
		Page:  pagingResult.Offset/pagingResult.Limit + 1,
	})
}

func (rpc *RpcController) AddTask(c *gin.Context) {
	var req model.AddTaskRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		rpc.Response(c, http.StatusBadRequest, 1, "bad request", nil)
		return
	}
	req.ScriptType = "python"
	err = rpc.TaskService.AddTask(req)
	if err != nil {
		rpc.Response(c, http.StatusConflict, 2, err.Error(), nil)
		return
	}

	rpc.ResponseOK(c, "")
}

func (rpc *RpcController) ListTasks(c *gin.Context) {
	list, err := rpc.TaskService.ListTask()
	if err != nil {
		rpc.Response(c, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	respList := []model.TaskListResponseItem{}
	for _, v := range list {
		respList = append(respList, convertTask(v))
	}
	rpc.ResponsePaging(c, model.PagingResult{
		Offset: 0,
		Limit:  10,
		Total:  0,
	}, nil, respList)
}

func convertTask(v db.Task) model.TaskListResponseItem {
	return model.TaskListResponseItem{
		BasicTask: model.BasicTask{
			Name:            v.Name,
			ScriptType:      v.ScriptType,
			ScriptFileName:  v.ScriptPath,
			Args:            v.Args,
			IntervalSeconds: v.IntervalSeconds,
		},
		BasicTaskResult: model.BasicTaskResult{
			Time:            SqlNullTimeToInt64Default(v.LastRunTime),
			Code:            v.LastRunCode,
			DurationSeconds: v.LastRunDuration,
			LogPath:         v.LastRunLogPath,
		},
	}
}
