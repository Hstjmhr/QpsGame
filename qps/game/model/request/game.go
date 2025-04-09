package request

import "game/conmponet/proto"

type RoomMessageReq struct {
	Type proto.RoomMessageType `json:"type"`
	Data RoomMessageData       `json:"data"`
}

type RoomMessageData struct {
	IsReady bool `json:"isReady"`
}
