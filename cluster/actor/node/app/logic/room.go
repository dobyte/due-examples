package logic

import (
	"context"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xconv"
	"github.com/dobyte/due/v2/utils/xtime"
	"sync"
	"time"
)

type Room struct {
	mgr     *RoomMgr
	id      int32
	name    string
	creator int64
	rw      sync.RWMutex
	members map[int64]time.Time
}

func NewRoom(mgr *RoomMgr, creator int64, name string) *Room {
	r := &Room{}
	r.mgr = mgr
	r.id = mgr.genRoomID()
	r.name = name
	r.creator = creator
	r.members = make(map[int64]time.Time)
	r.members[creator] = xtime.Now()

	return r
}

func (r *Room) ID() int32 {
	return r.id
}

func (r *Room) Name() string {
	return r.name
}

// 进入聊天室
func (r *Room) enterRoom(uid int64) (err error) {
	var ok bool

	r.rw.Lock()
	if _, ok = r.members[uid]; !ok {
		r.members[uid] = xtime.Now()
	}
	r.rw.Unlock()

	if ok {
		return errors.NewError(codes.IllegalRequest)
	}

	defer func() {
		if err != nil {
			r.rw.Lock()
			delete(r.members, uid)
			r.rw.Unlock()
		}
	}()

	if err = r.mgr.proxy.BindNode(context.Background(), uid); err != nil {
		log.Errorf("bind node failed: %v", err)
		return errors.NewError(err, codes.InternalError)
	}

	if err = r.mgr.proxy.BindActor(uid, roomActor, xconv.String(r.ID())); err != nil {
		log.Errorf("bind actor failed: %v", err)
		return errors.NewError(err, codes.InternalError)
	}

	return
}

func (r *Room) makeRoomInfo() *roomInfo {
	return &roomInfo{
		ID:      r.id,
		Name:    r.name,
		Creator: r.creator,
	}
}
