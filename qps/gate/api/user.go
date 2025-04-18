package api

import (
	"context"
	"framework/msError"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	common "msqp"
	"msqp/biz"
	"msqp/config"
	"msqp/jwts"
	"msqp/logs"
	"msqp/rpc"
	"time"
	"user/pb"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Register(ctx *gin.Context) {
	//接收参数
	var params pb.RegisterParams
	err2 := ctx.ShouldBindJSON(&params)
	if err2 != nil {
		logs.Error("Register request parse params err:%v", err2)
		common.Fail(ctx, biz.RequestDataError)
		return
	}

	response, err := rpc.UserClient.Register(context.TODO(), &params)
	if err != nil {
		common.Fail(ctx, msError.ToError(err))
		return
	}
	uid := response.Uid
	if len(uid) == 0 {
		common.Fail(ctx, biz.RequestDataError)
		return
	}
	logs.Info("uid:%s", uid)

	//gen token by uid jwt A.B.C
	//A:部分头(定义加密算法),B:部分存储数据,C:部分签名
	claims := jwts.CustomClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}
	token, err := jwts.GenToken(&claims, config.Conf.Jwt.Secret)
	if err != nil {
		logs.Error("Register jwt get token err:%v", err)
		common.Fail(ctx, biz.Fail)
		return
	}
	result := map[string]any{
		"token": token,
		"serverInfo": map[string]any{
			"host": config.Conf.Services["connector"].ClientHost,
			"port": config.Conf.Services["connector"].ClientPort,
		},
	}

	common.Success(ctx, result)
}
