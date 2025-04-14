package logic

type createRoomReq struct {
	Name string `json:"name"` // 房间名称
}

type createRoomRes struct {
	Code int                `json:"code"` // 响应码
	Data *createRoomResData `json:"data"` // 响应数据
}

type createRoomResData struct {
	Room *roomInfo `json:"room"`
}

type roomInfo struct {
	ID      int32  `json:"id"`      // 房间ID
	Name    string `json:"name"`    // 房间名称
	Creator int64  `json:"creator"` // 房间创建者
}

type enterRoomReq struct {
	ID int32 `json:"id"` // 房间ID
}

type enterRoomRes struct {
	Code int `json:"code"` // 响应码
}
