package user

import (
	"context"
	"due-examples/cluster/internal/codes"
	"due-examples/cluster/internal/service/grpc/user/pb"
	"github.com/dobyte/due/v2/cluster/mesh"
	"github.com/dobyte/due/v2/errors"
	"sync"
	"sync/atomic"
)

const service = "user"

type user struct {
	id       int64  // ID
	account  string // 账号
	password string // 密码
	nickname string // 昵称
}

type Server struct {
	pb.UnimplementedUserServer
	proxy    *mesh.Proxy
	uid      int64
	rw       sync.RWMutex
	users    map[int64]*user
	accounts map[string]*user
}

var _ pb.UserServer = &Server{}

func NewServer(proxy *mesh.Proxy) *Server {
	return &Server{
		proxy:    proxy,
		users:    make(map[int64]*user),
		accounts: make(map[string]*user),
	}
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider(service, &pb.User_ServiceDesc, s)
}

// Register 注册账号
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	s.rw.RLock()
	_, ok := s.accounts[req.Account]
	s.rw.RUnlock()
	if ok {
		return nil, errors.NewError(codes.AccountExists)
	}

	s.rw.Lock()
	defer s.rw.Unlock()

	if _, ok = s.accounts[req.Account]; ok {
		return nil, errors.NewError(codes.AccountExists)
	}

	uu := &user{
		id:       atomic.AddInt64(&s.uid, 1),
		account:  req.Account,
		password: req.Password,
		nickname: req.Nickname,
	}

	s.users[uu.id] = uu
	s.accounts[req.Account] = uu

	return nil, nil
}

// Login 登录账号
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	s.rw.RLock()
	uu, ok := s.accounts[req.Account]
	s.rw.RUnlock()

	if !ok {
		return nil, errors.NewError(codes.AccountNotExists)
	}

	if uu.password != req.Password {
		return nil, errors.NewError(codes.IncorrectAccountOrPassword)
	}

	return &pb.LoginReply{UID: uu.id}, nil
}

// FetchProfile 拉取资料
func (s *Server) FetchProfile(ctx context.Context, req *pb.FetchProfileRequest) (*pb.FetchProfileReply, error) {
	s.rw.RLock()
	uu, ok := s.users[req.UID]
	s.rw.RUnlock()

	if !ok {
		return nil, errors.NewError(codes.AccountNotExists)
	}

	return &pb.FetchProfileReply{
		Account:  uu.account,
		Nickname: uu.nickname,
	}, nil
}
