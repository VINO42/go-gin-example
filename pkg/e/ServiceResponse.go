package e

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServiceResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *ServiceResponse) Response(c *gin.Context, code int, msg string, data ...interface{}) {
	c.JSON(http.StatusOK, &ServiceResponse{
		Code: code,
		Msg:  msg,
		Data: data[0],
	})
}
func (r *ServiceResponse) Success(c gin.Context, msg string, data ...interface{}) {
	c.JSON(http.StatusOK, &ServiceResponse{
		Code: SUCCESS,
		Msg:  msg,
		Data: data[0],
	})
}

func (r *ServiceResponse) SuccessData(c *gin.Context, data ...interface{}) {

	c.JSON(http.StatusOK, &ServiceResponse{
		Code: SUCCESS,
		Msg:  GetMsg(SUCCESS),
		Data: data[0],
	})
}

func (r *ServiceResponse) Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &ServiceResponse{
		Code: ERROR,
		Msg:  msg,
	})
}

func (r *ServiceResponse) ErrorWithMsg(c *gin.Context) {
	c.JSON(http.StatusOK, &ServiceResponse{
		Code: ERROR,
		Msg:  GetMsg(ERROR),
	})
}
