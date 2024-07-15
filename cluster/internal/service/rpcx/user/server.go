package user

import (
	"context"
	"due-examples/cluster/internal/codes"
	"due-examples/cluster/internal/service/rpcx/user/pb"
	"github.com/dobyte/due/v2/cluster/mesh"
	"github.com/dobyte/due/v2/errors"
	"sync"
	"sync/atomic"
)

const (
	service     = "user" // 用于客户端定位服务，例如discovery://user
	servicePath = "User" // 服务路径要与pb中的服务路径保持一致
)

type user struct {
	id       int64  // ID
	account  string // 账号
	password string // 密码
	nickname string // 昵称
}

type User struct {
	proxy    *mesh.Proxy
	uid      int64
	rw       sync.RWMutex
	users    map[int64]*user
	accounts map[string]*user
}

var _ pb.UserAble = &User{}

func NewServer(proxy *mesh.Proxy) *User {
	return &User{
		proxy:    proxy,
		users:    make(map[int64]*user),
		accounts: make(map[string]*user),
	}
}

func (u *User) Init() {
	u.proxy.AddServiceProvider(service, servicePath, u)
}

// Register 注册账号
func (u *User) Register(ctx context.Context, req *pb.RegisterRequest, reply *pb.RegisterReply) error {
	u.rw.RLock()
	_, ok := u.accounts[req.Account]
	u.rw.RUnlock()
	if ok {
		return errors.NewError(codes.AccountExists)
	}

	u.rw.Lock()
	defer u.rw.Unlock()

	if _, ok = u.accounts[req.Account]; ok {
		return errors.NewError(codes.AccountExists)
	}

	uu := &user{
		id:       atomic.AddInt64(&u.uid, 1),
		account:  req.Account,
		password: req.Password,
		nickname: req.Nickname,
	}

	u.users[uu.id] = uu
	u.accounts[req.Account] = uu

	return nil
}

// Login 登录账号
func (u *User) Login(ctx context.Context, req *pb.LoginRequest, reply *pb.LoginReply) error {
	u.rw.RLock()
	uu, ok := u.accounts[req.Account]
	u.rw.RUnlock()

	if !ok {
		return errors.NewError(codes.AccountNotExists)
	}

	if uu.password != req.Password {
		return errors.NewError(codes.IncorrectAccountOrPassword)
	}

	reply.UID = uu.id

	return nil
}

// FetchProfile 拉取资料
func (u *User) FetchProfile(ctx context.Context, req *pb.FetchProfileRequest, reply *pb.FetchProfileReply) error {
	u.rw.RLock()
	uu, ok := u.users[req.UID]
	u.rw.RUnlock()

	if !ok {
		return errors.NewError(codes.AccountNotExists)
	}

	reply.Account = uu.account
	reply.Nickname = uu.nickname

	return nil
}
