package logic

import (
	"core/models/entity"
	"fmt"
	"framework/msError"
	"framework/remote"
	"game/conmponet/room"
	"math/rand"
	"msqp/biz"
	"sync"
	"time"
)

type UnionManager struct {
	sync.RWMutex
	unionList map[int64]*Union
}

func NewUnionManager() *UnionManager {
	return &UnionManager{
		unionList: make(map[int64]*Union),
	}
}

func (u *UnionManager) GetUnion(unionId int64) *Union {
	//使用互斥锁（Mutex）来保护 unionList 的并发访问。u.Lock() 将锁定 UnionManager 实例，
	//确保在执行方法期间，其他 goroutine 不能修改 unionList。
	//defer u.Unlock() 确保在方法结束时自动释放锁，无论是正常返回还是因错误退出。
	u.Lock()
	defer u.Unlock()
	union, ok := u.unionList[unionId]
	if ok {
		return union
	}
	union = NewUnion(u)
	u.unionList[unionId] = union
	return union
}

func (u *UnionManager) CreateRoomId() string {
	//随机数的方式去创建
	roomId := u.genRoomId()
	for _, v := range u.unionList {
		_, ok := v.RoomList[roomId]
		if ok {
			return u.CreateRoomId()
		}
	}
	return roomId
}

func (u *UnionManager) genRoomId() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// 房间号是6位数
	roomIdInt := rand.Int63n(999999)
	if roomIdInt < 100000 {
		roomIdInt += 100000
	}
	return fmt.Sprintf("%d", roomIdInt)
}

func (u *UnionManager) GetRoomById(roomId string) *room.Room {
	for _, v := range u.unionList {
		r, ok := v.RoomList[roomId]
		if ok {
			return r
		}
	}
	return nil
}

func (u *UnionManager) JoinRoom(session *remote.Session, roomId string, data *entity.User) *msError.Error {
	for _, v := range u.unionList {
		r, ok := v.RoomList[roomId]
		if ok {
			return r.JoinRoom(session, data)
		}
	}
	return biz.RoomNotExist
}
