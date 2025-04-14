package logic

import (
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/utils/xconv"
	"sync"
	"sync/atomic"
)

type RoomMgr struct {
	id    atomic.Int32
	proxy *node.Proxy
	rw    sync.RWMutex
	rooms map[int32]*Room
}

func NewRoomMgr(proxy *node.Proxy) *RoomMgr {
	return &RoomMgr{proxy: proxy}
}

// 创建聊天室
func (mgr *RoomMgr) createRoom(creator int64, name string) (*Room, error) {
	room := NewRoom(mgr, creator, name)

	actor, err := mgr.proxy.Spawn(NewRoomProcessor, node.WithActorID(xconv.String(room.ID())), node.WithActorArgs(room))
	if err != nil {
		return nil, errors.NewError(err, codes.InternalError)
	}

	if err = mgr.proxy.BindActor(creator, roomActor, actor.ID()); err != nil {
		actor.Destroy()
		return nil, errors.NewError(err, codes.InternalError)
	}

	mgr.rw.Lock()
	mgr.rooms[room.ID()] = room
	mgr.rw.Unlock()

	return room, nil
}

// 获取聊天室
func (mgr *RoomMgr) getRoom(roomID int32) (*Room, bool) {
	mgr.rw.RLock()
	defer mgr.rw.RUnlock()

	room, ok := mgr.rooms[roomID]
	return room, ok
}

// 进入聊天室
func (mgr *RoomMgr) enterRoom(uid int64, roomID int32) error {
	room, ok := mgr.getRoom(roomID)
	if !ok {
		return errors.NewError(codes.NotFound)
	}

	return room.enterRoom(uid)
}

// 生成房间ID
func (mgr *RoomMgr) genRoomID() int32 {
	return mgr.id.Add(1)
}
