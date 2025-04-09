package room

import (
	"framework/remote"
	"game/conmponet/proto"
)

type GameFrame interface {
	GetGameData(session *remote.Session) any
	StartGame(session *remote.Session, user *proto.RoomUser)
	GameMessageHandle(user *proto.RoomUser, session *remote.Session, msg []byte)
}
