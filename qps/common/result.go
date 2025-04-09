package common

import (
	"framework/msError"
	"github.com/gin-gonic/gin"
	"msqp/biz"
	"net/http"
)

type Result struct {
	Code int `json:"code"`
	Msg  any `json:"msg"`
}

func F(err *msError.Error) any {
	return Result{
		Code: err.Code,
	}
}

func S(data any) Result {
	return Result{
		Code: biz.OK,
		Msg:  data,
	}
}

func Fail(ctx *gin.Context, err *msError.Error) {
	ctx.JSON(http.StatusOK, Result{
		Code: err.Code,
		Msg:  err.Err.Error(),
	})
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Result{
		Code: biz.OK,
		Msg:  data,
	})
}
