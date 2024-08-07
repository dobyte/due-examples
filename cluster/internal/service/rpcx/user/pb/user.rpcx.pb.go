// Code generated by protoc-gen-rpcx. DO NOT EDIT.
// versions:
// - protoc-gen-rpcx v0.3.0
// - protoc          v5.26.1
// source: user.proto

package pb

import (
	context "context"
	client "github.com/smallnest/rpcx/client"
	protocol "github.com/smallnest/rpcx/protocol"
	server "github.com/smallnest/rpcx/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = context.TODO
var _ = server.NewServer
var _ = client.NewClient
var _ = protocol.NewMessage

// ================== interface skeleton ===================
type UserAble interface {
	// UserAble can be used for interface verification.

	// Register is server rpc method as defined
	Register(ctx context.Context, args *RegisterRequest, reply *RegisterReply) (err error)

	// Login is server rpc method as defined
	Login(ctx context.Context, args *LoginRequest, reply *LoginReply) (err error)

	// FetchProfile is server rpc method as defined
	FetchProfile(ctx context.Context, args *FetchProfileRequest, reply *FetchProfileReply) (err error)
}

// ================== server skeleton ===================
type UserImpl struct{}

// ServeForUser starts a server only registers one service.
// You can register more services and only start one server.
// It blocks until the application exits.
func ServeForUser(addr string) error {
	s := server.NewServer()
	s.RegisterName("User", new(UserImpl), "")
	return s.Serve("tcp", addr)
}

// Register is server rpc method as defined
func (s *UserImpl) Register(ctx context.Context, args *RegisterRequest, reply *RegisterReply) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = RegisterReply{}

	return nil
}

// Login is server rpc method as defined
func (s *UserImpl) Login(ctx context.Context, args *LoginRequest, reply *LoginReply) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = LoginReply{}

	return nil
}

// FetchProfile is server rpc method as defined
func (s *UserImpl) FetchProfile(ctx context.Context, args *FetchProfileRequest, reply *FetchProfileReply) (err error) {
	// TODO: add business logics

	// TODO: setting return values
	*reply = FetchProfileReply{}

	return nil
}

// ================== client stub ===================
// User is a client wrapped XClient.
type UserClient struct {
	xclient client.XClient
}

// NewUserClient wraps a XClient as UserClient.
// You can pass a shared XClient object created by NewXClientForUser.
func NewUserClient(xclient client.XClient) *UserClient {
	return &UserClient{xclient: xclient}
}

// NewXClientForUser creates a XClient.
// You can configure this client with more options such as etcd registry, serialize type, select algorithm and fail mode.
func NewXClientForUser(addr string) (client.XClient, error) {
	d, err := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	if err != nil {
		return nil, err
	}

	opt := client.DefaultOption
	opt.SerializeType = protocol.ProtoBuffer

	xclient := client.NewXClient("User", client.Failtry, client.RoundRobin, d, opt)

	return xclient, nil
}

// Register is client rpc method as defined
func (c *UserClient) Register(ctx context.Context, args *RegisterRequest) (reply *RegisterReply, err error) {
	reply = &RegisterReply{}
	err = c.xclient.Call(ctx, "Register", args, reply)
	return reply, err
}

// Login is client rpc method as defined
func (c *UserClient) Login(ctx context.Context, args *LoginRequest) (reply *LoginReply, err error) {
	reply = &LoginReply{}
	err = c.xclient.Call(ctx, "Login", args, reply)
	return reply, err
}

// FetchProfile is client rpc method as defined
func (c *UserClient) FetchProfile(ctx context.Context, args *FetchProfileRequest) (reply *FetchProfileReply, err error) {
	reply = &FetchProfileReply{}
	err = c.xclient.Call(ctx, "FetchProfile", args, reply)
	return reply, err
}

// ================== oneclient stub ===================
// UserOneClient is a client wrapped oneClient.
type UserOneClient struct {
	serviceName string
	oneclient   *client.OneClient
}

// NewUserOneClient wraps a OneClient as UserOneClient.
// You can pass a shared OneClient object created by NewOneClientForUser.
func NewUserOneClient(oneclient *client.OneClient) *UserOneClient {
	return &UserOneClient{
		serviceName: "User",
		oneclient:   oneclient,
	}
}

// ======================================================

// Register is client rpc method as defined
func (c *UserOneClient) Register(ctx context.Context, args *RegisterRequest) (reply *RegisterReply, err error) {
	reply = &RegisterReply{}
	err = c.oneclient.Call(ctx, c.serviceName, "Register", args, reply)
	return reply, err
}

// Login is client rpc method as defined
func (c *UserOneClient) Login(ctx context.Context, args *LoginRequest) (reply *LoginReply, err error) {
	reply = &LoginReply{}
	err = c.oneclient.Call(ctx, c.serviceName, "Login", args, reply)
	return reply, err
}

// FetchProfile is client rpc method as defined
func (c *UserOneClient) FetchProfile(ctx context.Context, args *FetchProfileRequest) (reply *FetchProfileReply, err error) {
	reply = &FetchProfileReply{}
	err = c.oneclient.Call(ctx, c.serviceName, "FetchProfile", args, reply)
	return reply, err
}
