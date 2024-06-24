package logic

type room struct {
}

func newRoom() *room {

}

// Kind Actor类型
func (r *room) Kind() string {
	return "room"
}

// Dispatch 派发消息
func (r *room) Dispatch() {
}
