package handler

import (
	"context"
	"core/repo"
	"core/service"
	"encoding/json"
	"framework/remote"
	"game/logic"
	"game/model/request"
	common "msqp"
	"msqp/biz"
)

type UnionHandler struct {
	um          *logic.UnionManager
	userService *service.UserService
}

func (h *UnionHandler) CreatRoom(session *remote.Session, msg []byte) any {
	//union 联盟 持有房间
	//unionManager 管理联盟
	//room 房间又关联 game接口 实现多个不同的游戏
	//1.接收参数
	uid := session.GetUid()
	if len(uid) <= 0 {
		return common.F(biz.InvalidUsers)
	}
	var req request.CreateRoomReq
	if err := json.Unmarshal(msg, &req); err != nil {
		return common.F(biz.RequestDataError)
	}
	//2.根据session 用户id 查询用户的信息
	userData, err := h.userService.FindUserByUid(context.TODO(), uid)
	if err != nil {
		return common.F(err)
	}
	if userData == nil {
		return common.F(biz.InvalidUsers)
	}
	//3.根据游戏规则 游戏类型 用户信息（创建房间的用户）
	// TODO 需要判断session中是否已经有roomID，如果有，就已经代表此用户已经在房间中了，就不能再次创建房间了
	union := h.um.GetUnion(req.UnionID)
	err = union.CreateRoom(h.userService, session, req, userData)
	if err != nil {
		return common.F(biz.InvalidUsers)
	}
	return common.S(nil)
}

func (h *UnionHandler) JoinRoom(session *remote.Session, msg []byte) any {
	uid := session.GetUid()
	if len(uid) <= 0 {
		return common.F(biz.InvalidUsers)
	}
	var req request.JoinRoomReq
	if err := json.Unmarshal(msg, &req); err != nil {
		return common.F(biz.RequestDataError)
	}
	userData, err := h.userService.FindUserByUid(context.TODO(), uid)
	if err != nil {
		return common.F(err)
	}
	if userData == nil {
		return common.F(biz.InvalidUsers)
	}
	bizErr := h.um.JoinRoom(session, req.RoomID, userData)
	if bizErr != nil {
		return common.F(bizErr)
	}
	return common.S(nil)

}

func NewUnionHandler(r *repo.Manager, um *logic.UnionManager) *UnionHandler {
	return &UnionHandler{
		um:          um,
		userService: service.NewUserService(r),
	}
}
