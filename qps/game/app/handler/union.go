package handler

import (
	"core/repo"
	"core/service"
	"encoding/json"
	"framework/remote"
	"hall/model/request"
	"hall/model/response"
	common "msqp"
	"msqp/biz"
	"msqp/logs"
)

type UserHandler struct {
	userService *service.UserService
}

func (h *UserHandler) UpdateUserAddress(session *remote.Session, msg []byte) any {
	logs.Info("UpdateUserAddress msg:%v", string(msg))
	var req request.UpdateUserAddressReq
	if err := json.Unmarshal(msg, &req); err != nil {
		return common.F(biz.RequestDataError)
	}
	err := h.userService.UpdateUserAddressByUid(session.GetUid(), req)
	if err != nil {
		return common.F(biz.SqlError)
	}
	res := response.UpdateUserAddressRes{}
	res.Code = biz.OK
	res.UpdateUserData = req
	return res
}

func NewUserHandler(r *repo.Manager) *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(r),
	}
}
