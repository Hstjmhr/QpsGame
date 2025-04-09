package route

import (
	"core/repo"
	"framework/node"
	"game/handler"
	"game/logic"
)

func Register(r *repo.Manager) node.LogicHandler {
	handlers := make(node.LogicHandler)
	um := logic.NewUnionManager()
	UnionHandler := handler.NewUnionHandler(r, um)
	handlers["unionHandler.createRoom"] = UnionHandler.CreatRoom
	handlers["unionHandler.joinRoom"] = UnionHandler.JoinRoom
	GameHandler := handler.NewGameHandler(r, um)
	handlers["gameHandler.roomMessageNotify"] = GameHandler.RoomMessageNotify
	handlers["gameHandler.gameMessageNotify"] = GameHandler.GameMessageNotify
	return handlers
}
