package request

import "game/conmponet/proto"

type CreateRoomReq struct {
	UnionID     int64          `json:"unionID"`
	GameRulerID string         `json:"gameRulerID"`
	GameRule    proto.GameRule `json:"gameRule"`
}

type JoinRoomReq struct {
	RoomID string `json:"roomID"`
}
