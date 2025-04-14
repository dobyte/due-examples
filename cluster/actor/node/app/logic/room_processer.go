package logic

import "github.com/dobyte/due/v2/cluster/node"

const roomActor = "room"

type RoomProcessor struct {
	node.BaseProcessor
	actor *node.Actor
	room  *Room
}

func NewRoomProcessor(actor *node.Actor, args ...any) node.Processor {
	return &RoomProcessor{
		actor: actor,
		room:  args[0].(*Room),
	}
}

// Kind 设置处理器类型
func (r *RoomProcessor) Kind() string {
	return roomActor
}

// Init 初始化处理器
func (r *RoomProcessor) Init() {

}
