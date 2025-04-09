package base

import (
	"framework/remote"
	"game/conmponet/proto"
)

type RoomFrame interface {
	GetUsers() map[string]*proto.RoomUser
	GetId() string
	EndGame(session *remote.Session)
	UserReady(uid string, session *remote.Session)
}
