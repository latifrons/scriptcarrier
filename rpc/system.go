package rpc

import (
	"github.com/gin-gonic/gin"
)

func (rpc *RpcController) Health(c *gin.Context) {
	rpc.ResponseOK(c, "okk")
}
