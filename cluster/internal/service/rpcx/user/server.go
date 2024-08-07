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

type Server struct {
	proxy    *mesh.Proxy
	uid      int64
	rw       sync.RWMutex
	users    map[int64]*user
	accounts map[string]*user
}

var _ pb.UserAble = &Server{}

func NewServer(proxy *mesh.Proxy) *Server {
	return &Server{
		proxy:    proxy,
		users:    make(map[int64]*user),
		accounts: make(map[string]*user),
	}
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider(service, servicePath, s)
}

// Register 注册账号
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest, reply *pb.RegisterReply) error {
	s.rw.RLock()
	_, ok := s.accounts[req.Account]
	s.rw.RUnlock()
	if ok {
		return errors.NewError(codes.AccountExists)
	}

	s.rw.Lock()
	defer s.rw.Unlock()

	if _, ok = s.accounts[req.Account]; ok {
		return errors.NewError(codes.AccountExists)
	}

	uu := &user{
		id:       atomic.AddInt64(&s.uid, 1),
		account:  req.Account,
		password: req.Password,
		nickname: req.Nickname,
	}

	s.users[uu.id] = uu
	s.accounts[req.Account] = uu

	return nil
}

// Login 登录账号
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest, reply *pb.LoginReply) error {
	s.rw.RLock()
	uu, ok := s.accounts[req.Account]
	s.rw.RUnlock()

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
func (s *Server) FetchProfile(ctx context.Context, req *pb.FetchProfileRequest, reply *pb.FetchProfileReply) error {
	s.rw.RLock()
	uu, ok := s.users[req.UID]
	s.rw.RUnlock()

	if !ok {
		return errors.NewError(codes.AccountNotExists)
	}

	reply.Account = uu.account
	reply.Nickname = uu.nickname

	return nil
}
