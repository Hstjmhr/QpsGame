package sz

import (
	"encoding/json"
	"fmt"
	"framework/remote"
	"game/conmponet/base"
	"game/conmponet/proto"
	"github.com/jinzhu/copier"
	"msqp/logs"
	"msqp/utils"
	"time"
)

type GameFrame struct {
	r          base.RoomFrame
	gameRule   proto.GameRule
	gameData   *GameData
	Logic      *Logic
	GameResult *GameResult
}

func (g *GameFrame) GameMessageHandle(user *proto.RoomUser, session *remote.Session, msg []byte) {
	//1. 解析参数
	var req MessageReq
	json.Unmarshal(msg, &req)
	//2. 根据不同的类型触发不同的操作
	if req.Type == GameLookNotify {
		g.onGameLook(user, session, req.Data.Cuopai)
	} else if req.Type == GamePourScoreNotify {
		g.onGamePourScore(user, session, req.Data.Score, req.Data.Type)
	} else if req.Type == GameCompareNotify {
		g.onGameCompare(user, session, req.Data.ChairID)
	} else if req.Type == GameAbandonNotify {
		g.onGameAbandon(user, session)
	}
}

func NewGameFrame(rule proto.GameRule, r base.RoomFrame) *GameFrame {
	gameData := initGameData(rule)
	return &GameFrame{
		r:        r,
		gameRule: rule,
		gameData: gameData,
		Logic:    NewLogic(),
	}
}

func (g *GameFrame) ServerMessagePush(users []string, data any, session *remote.Session) {
	session.Push(users, data, "ServerMessagePush")
}

func initGameData(rule proto.GameRule) *GameData {
	g := &GameData{
		GameType:   GameType(rule.GameFrameType),
		BaseScore:  rule.BaseScore,
		ChairCount: rule.MaxPlayerCount,
	}
	g.PourScores = make([][]int, g.ChairCount)
	g.HandCards = make([][]int, g.ChairCount)
	g.LookCards = make([]int, g.ChairCount)
	g.CurScores = make([]int, g.ChairCount)
	g.UserStatusArray = make([]UserStatus, g.ChairCount)
	g.UserTrustArray = []bool{false, false, false, false, false, false, false, false, false, false}
	g.Loser = make([]int, 0)
	g.Winner = make([]int, 0)
	return g
}

func (g *GameFrame) GetGameData(session *remote.Session) any {
	user := g.r.GetUsers()[session.GetUid()]
	//判断当前用户 是否已经看牌 如果已经看牌 返回牌 但是对其他用 户仍旧是隐藏状态
	//深拷贝
	var gameData GameData
	copier.CopyWithOption(&gameData, g.gameData, copier.Option{DeepCopy: true})
	for i := 0; i < g.gameData.ChairCount; i++ {
		if g.gameData.HandCards[i] != nil {
			gameData.HandCards[i] = make([]int, 3)
		} else {
			gameData.HandCards[i] = nil
		}
	}
	if g.gameData.LookCards[user.ChairID] == 1 {
		// 已经看牌了
		gameData.HandCards[user.ChairID] = g.gameData.HandCards[user.ChairID]
	}
	return gameData
}

func (g *GameFrame) StartGame(session *remote.Session, user *proto.RoomUser) {
	//1.用户信息变更推送
	users := g.getAllUsers()
	fmt.Println(users)
	g.ServerMessagePush(users, UpdateUserInfoPushGold(user.UserInfo.Gold), session)
	//2.庄家推送
	if g.gameData.CurBureau == 0 {
		//庄家是每次开始游戏 首次进行操作的座次
		g.gameData.BankerChairID = utils.Rand(len(users))
	}
	g.ServerMessagePush(users, GameBankerPushData(g.gameData.BankerChairID), session)
	//3.局数推送
	g.gameData.CurBureau++
	g.ServerMessagePush(users, GameBureauPushData(g.gameData.CurBureau), session)
	//4.游戏状态推送，分两步 第一步推送发牌，第二部推送下分
	g.gameData.GameStatus = SendCards
	g.ServerMessagePush(users, GameStatusPushData(g.gameData.GameStatus, 0), session)
	//5.发牌推送
	g.sendCards(session)
	//6.下分推送
	g.gameData.GameStatus = PourScore
	g.ServerMessagePush(users, GameStatusPushData(g.gameData.GameStatus, 30), session)
	g.gameData.CurScore = g.gameRule.AddScores[0] * g.gameRule.BaseScore
	for _, v := range g.r.GetUsers() {
		g.ServerMessagePush([]string{v.UserInfo.Uid}, GamePourScorePushData(v.ChairID, g.gameData.CurScore, g.gameData.CurScore, 1, 0), session)
	}
	//7. 轮数推送
	g.gameData.Round = 1
	g.ServerMessagePush(users, GameRoundPushData(g.gameData.Round), session)
	//8. 操作推送
	for _, v := range g.r.GetUsers() {
		g.ServerMessagePush([]string{v.UserInfo.Uid}, GameTurnPushData(g.gameData.CurChairID, g.gameData.CurScore), session)
	}

}

func (g *GameFrame) getAllUsers() []string {
	users := make([]string, 0)
	for _, v := range g.r.GetUsers() {
		users = append(users, v.UserInfo.Uid)
	}
	return users
}

func (g *GameFrame) sendCards(session *remote.Session) {
	//这里开始发牌 牌相关的逻辑
	//1.洗牌 然后发牌
	g.Logic.washCards()
	for i := 0; i < g.gameData.ChairCount; i++ {
		if g.IsPlayingChairID(i) {
			g.gameData.HandCards[i] = g.Logic.getCards()
		}
	}
	//发牌后 推送的时候 如果没有看牌的话 暗牌
	hands := make([][]int, g.gameData.ChairCount)
	for i, v := range g.gameData.HandCards {
		if v != nil {
			hands[i] = []int{0, 0, 0}
		}
	}
	g.ServerMessagePush(g.getAllUsers(), GameSendCardsPushData(hands), session)
}

func (g *GameFrame) IsPlayingChairID(chairID int) bool {
	for _, v := range g.r.GetUsers() {
		if v.ChairID == chairID && v.UserStatus == proto.Playing {
			return true
		}
	}
	return false
}

func (g *GameFrame) onGameLook(user *proto.RoomUser, session *remote.Session, cuopai bool) {
	//判断 如果是当前用户 推送其牌 给其他用户 推送此用户已经看牌
	if g.gameData.GameStatus != PourScore {
		logs.Warn("ID:%s room,sanzhang game look err:gameStatus=%d,curChairID=%d,chairID=%d",
			g.r.GetId(), g.gameData.GameStatus, g.gameData.CurChairID, user.ChairID)
		return
	}
	if !g.IsPlayingChairID(user.ChairID) {
		logs.Warn("ID:%s room,sanzhang game look err: not playing",
			g.r.GetId())
		return
	}
	//代表已看牌
	g.gameData.UserStatusArray[user.ChairID] = Look
	g.gameData.LookCards[user.ChairID] = 1
	for _, v := range g.r.GetUsers() {
		if g.gameData.CurChairID == v.ChairID {
			//代表操作用户
			g.ServerMessagePush([]string{v.UserInfo.Uid}, GameLookPushData(g.gameData.CurChairID, g.gameData.HandCards[v.ChairID], cuopai), session)
		} else {
			g.ServerMessagePush([]string{v.UserInfo.Uid}, GameLookPushData(g.gameData.CurChairID, nil, cuopai), session)
		}
	}
}

func (g *GameFrame) onGamePourScore(user *proto.RoomUser, session *remote.Session, score int, t int) {
	//1. 处理下分 保存用户下的分数 同时推送当前用户下分的信息到客户端
	if g.gameData.GameStatus != PourScore || g.gameData.CurChairID != user.ChairID {
		logs.Warn("ID:%s room,sanzhang onGamePourScore err:gameStatus=%d,curChairID=%d,chairID=%d",
			g.r.GetId(), g.gameData.GameStatus, g.gameData.CurChairID, user.ChairID)
		return
	}
	if !g.IsPlayingChairID(user.ChairID) {
		logs.Warn("ID:%s room,sanzhang onGamePourScore err: not playing",
			g.r.GetId())
		return
	}
	if score < 0 {
		logs.Warn("ID:%s room,sanzhang onGamePourScore err: score lt zero",
			g.r.GetId())
		return
	}
	if g.gameData.PourScores[user.ChairID] == nil {
		g.gameData.PourScores[user.ChairID] = make([]int, 0)
	}
	g.gameData.PourScores[user.ChairID] = append(g.gameData.PourScores[user.ChairID], score)
	// 所有人的分数
	scores := 0
	for i := 0; i < g.gameData.ChairCount; i++ {
		if g.gameData.PourScores[i] != nil {
			for _, sc := range g.gameData.PourScores[i] {
				scores += sc
			}
		}
	}
	//  当前座次的总分
	chairCount := 0
	for _, sc := range g.gameData.PourScores[user.ChairID] {
		chairCount += sc
	}
	g.ServerMessagePush(g.getAllUsers(), GamePourScorePushData(user.ChairID, score, chairCount, scores, t), session)
	//2. 结束下分 座次移动到下一位 推送轮次 推送游戏状态 推送操作的座次
	g.endPourScore(session)
}

func (g *GameFrame) endPourScore(session *remote.Session) {
	//1. 推送轮次 TODO 轮数大于规则的限制 结束游戏 进行结算
	round := g.getCurRound()
	g.ServerMessagePush(g.getAllUsers(), GameRoundPushData(round), session)
	// 判断当前的玩家没有lose的 只剩下一个的时候
	gamerCount := 0
	for i := 0; i < g.gameData.ChairCount; i++ {
		if g.IsPlayingChairID(i) && !utils.Contains(g.gameData.Loser, i) {
			gamerCount++
		}
	}
	if gamerCount == 1 {
		g.StartResult(session)
	} else {
		//2. 座次向前移动一位
		for i := 0; i < g.gameData.ChairCount; i++ {
			g.gameData.CurChairID++
			g.gameData.CurChairID = g.gameData.CurChairID % g.gameData.ChairCount
			if g.IsPlayingChairID(g.gameData.CurChairID) {
				break
			}
		}
		//推送游戏状态
		g.gameData.GameStatus = PourScore
		g.ServerMessagePush(g.getAllUsers(), GameStatusPushData(g.gameData.GameStatus, 30), session)
		//该谁操作了
		g.ServerMessagePush(g.getAllUsers(), GameTurnPushData(g.gameData.CurChairID, g.gameData.CurScore), session)
	}
}

func (g *GameFrame) getCurRound() int {
	cur := g.gameData.CurChairID
	for i := 0; i < g.gameData.CurChairID; i++ {
		cur++
		cur = cur % g.gameData.CurChairID
		if g.IsPlayingChairID(cur) {
			return len(g.gameData.PourScores[cur])
		}
	}
	return 1
}

func (g *GameFrame) onGameCompare(user *proto.RoomUser, session *remote.Session, chairID int) {
	//1. TODO 先下分 跟注结束之后 进行比牌
	//2. 比牌
	fromChairID := user.ChairID
	toChairID := chairID
	result := g.Logic.CompareCards(g.gameData.HandCards[fromChairID], g.gameData.HandCards[toChairID])
	//3. 处理比牌结果 推送轮数 状态 显示结果等信息
	if result == 0 {
		//主动比牌者 如果结果是和 比牌者输
		result = -1
	}
	winChairID := -1
	loseChairID := -1
	if result > 0 {
		//发起比牌者赢，其他玩家输
		g.ServerMessagePush(g.getAllUsers(), GameComparePushData(fromChairID, toChairID, fromChairID, toChairID), session)
		winChairID = fromChairID
		loseChairID = toChairID
	} else if result < 0 {
		//玩家赢，发起比牌者输
		g.ServerMessagePush(g.getAllUsers(), GameComparePushData(fromChairID, toChairID, toChairID, fromChairID), session)
		winChairID = toChairID
		loseChairID = fromChairID
	}
	if winChairID != -1 && loseChairID != -1 {
		g.gameData.UserStatusArray[winChairID] = Win
		g.gameData.UserStatusArray[loseChairID] = Lose
		g.gameData.Winner = append(g.gameData.Winner, winChairID)
		g.gameData.Loser = append(g.gameData.Loser, loseChairID)
	}
	//TODO 赢了之后 继续和其他人进行比牌
	if winChairID == fromChairID {
	}
	g.endPourScore(session)
}

func (g *GameFrame) StartResult(session *remote.Session) {
	//推送 游戏结果状态
	g.gameData.GameStatus = Result
	g.ServerMessagePush(g.getAllUsers(), GameStatusPushData(g.gameData.GameStatus, 0), session)
	if g.GameResult == nil {
		g.GameResult = new(GameResult)
	}
	g.GameResult.Winners = g.gameData.Winner
	g.GameResult.HandCards = g.gameData.HandCards
	g.GameResult.CurScores = g.gameData.CurScores
	g.GameResult.Losers = g.gameData.Loser
	WinScore := make([]int, g.gameData.ChairCount)
	for i := range WinScore {
		if g.gameData.PourScores != nil {
			scores := 0
			for _, v := range g.gameData.PourScores[i] {
				scores += v
			}
			WinScore[i] = -scores
			for win := range g.gameData.Winner {
				WinScore[win] += scores / len(g.gameData.Winner)
			}
		}
	}
	g.GameResult.WinScores = WinScore
	g.ServerMessagePush(g.getAllUsers(), GameResultPushData(g.GameResult), session)
	//结算完成 重置游戏 开始下一把
	g.resetGame(session)
	g.gameEnd(session)
}

func (gf *GameFrame) resetGame(session *remote.Session) {
	g := &GameData{
		GameType:   GameType(gf.gameRule.GameFrameType),
		BaseScore:  gf.gameRule.BaseScore,
		ChairCount: gf.gameRule.MaxPlayerCount,
	}
	g.PourScores = make([][]int, g.ChairCount)
	g.HandCards = make([][]int, g.ChairCount)
	g.LookCards = make([]int, g.ChairCount)
	g.CurScores = make([]int, g.ChairCount)
	g.UserStatusArray = make([]UserStatus, g.ChairCount)
	g.UserTrustArray = []bool{false, false, false, false, false, false, false, false, false, false}
	g.Loser = make([]int, 0)
	g.Winner = make([]int, 0)
	g.GameStatus = GameStatus(None)
	gf.gameData = g
	gf.SendGameStatus(g.GameStatus, 0, session)
	gf.r.EndGame(session)
}

func (g *GameFrame) SendGameStatus(status GameStatus, tick int, session *remote.Session) {
	g.ServerMessagePush(g.getAllUsers(), GameStatusPushData(status, tick), session)
}

func (g *GameFrame) gameEnd(session *remote.Session) {
	//赢家当庄家
	for i := 0; i < g.gameData.ChairCount; i++ {
		if g.GameResult.WinScores[i] > 0 {
			g.gameData.BankerChairID = i
			g.gameData.CurChairID = g.gameData.BankerChairID
		}
	}
	time.AfterFunc(5*time.Second, func() {
		for _, v := range g.r.GetUsers() {
			g.r.UserReady(v.UserInfo.Uid, session)
		}
	})
}

func (g *GameFrame) onGameAbandon(user *proto.RoomUser, session *remote.Session) {
	if !g.IsPlayingChairID(user.ChairID) {
		return
	}
	if utils.Contains(g.gameData.Loser, user.ChairID) {
		return
	}
	g.gameData.Loser = append(g.gameData.Loser, user.ChairID)
	for i := 0; i < g.gameData.ChairCount; i++ {
		if g.IsPlayingChairID(i) && i != user.ChairID {
			g.gameData.Winner = append(g.gameData.Winner, g.gameData.BankerChairID)
		}
	}
	g.gameData.UserStatusArray[user.ChairID] = Abandon
	//推送弃牌状态
	g.send(GameAbandonPushData(user.ChairID, g.gameData.UserStatusArray[user.ChairID]), session)

	time.AfterFunc(2*time.Second, func() {
		g.endPourScore(session)
	})
}

func (g *GameFrame) send(data any, session *remote.Session) {
	g.ServerMessagePush(g.getAllUsers(), data, session)
}
