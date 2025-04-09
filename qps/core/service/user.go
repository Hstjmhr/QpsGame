package service

import (
	"connector/models/request"
	"context"
	"core/dao"
	"core/models/entity"
	"core/repo"
	"fmt"
	"framework/game"
	"framework/msError"
	hall "hall/model/request"
	"msqp"
	"msqp/biz"
	"msqp/logs"
	"time"
)

type UserService struct {
	UserDao *dao.UserDao
}

func (s *UserService) FindAndSavaUserByUid(ctx context.Context, uid string, info request.UserInfo) (*entity.User, error) {
	//查询mongo 有返回，没有新增
	user, err := s.UserDao.FindUserByUid(ctx, uid)
	if err != nil {
		logs.Error("[UserService] FindAndSavaUserByUid user err:%v", err)
		return nil, err
	}
	if user == nil {
		//save
		user = &entity.User{}
		user.Uid = uid
		user.Gold = int64(game.Conf.GameConfig["startGold"]["value"].(float64))
		user.Avatar = common.Default(info.Avatar, "Common/head_icon_default")
		//if len(info.Avatar) == 0 {
		//	user.Avatar = "Common/head_icon_default"
		//} else {
		//	user.Avatar = info.Avatar
		//}
		user.Nickname = common.Default(info.Nickname, fmt.Sprintf("%s%s", "新用户", uid))
		user.Sex = info.Sex // 0男1女
		user.CreateTime = time.Now().UnixMilli()
		user.LastLoginTime = time.Now().UnixMilli()
		err = s.UserDao.Insert(context.TODO(), user)
		if err != nil {
			logs.Error("[UserService] FindAndSavaUserByUid insert user err:%v", err)
			return nil, err
		}
	}
	return user, nil
}

func (s *UserService) FindUserByUid(ctx context.Context, uid string) (*entity.User, *msError.Error) {
	//查询mongo 有返回，没有新增
	user, err := s.UserDao.FindUserByUid(ctx, uid)
	if err != nil {
		logs.Error("[UserService] FindUserByUid user err:%v", err)
		return nil, biz.SqlError
	}

	return user, nil
}

func (s *UserService) UpdateUserAddressByUid(uid string, req hall.UpdateUserAddressReq) error {
	user := &entity.User{
		Uid:      uid,
		Address:  req.Address,
		Location: req.Location,
	}
	err := s.UserDao.UpdateUserAddressByUid(context.TODO(), user)
	if err != nil {
		logs.Error("userDao.UpdateUserAddressByUid err%v", err)
		return err
	}
	return nil
}

func NewUserService(r *repo.Manager) *UserService {
	return &UserService{
		UserDao: dao.NewUserDao(r),
	}
}
