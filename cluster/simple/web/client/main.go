package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	duehttp "github.com/dobyte/due/component/http/v2"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xtime"
)

const url = "http://127.0.0.1:8080/greet"

func main() {
	for {
		sendRequest()

		time.Sleep(time.Second)
	}
}

type greetReq struct {
	Message string `json:"message"`
}

// 响应
type greetRes struct {
	Message string `json:"message"`
}

func sendRequest() {
	b, err := json.Marshal(&greetReq{
		Message: fmt.Sprintf("I'm ws client, and the current time is: %s", xtime.Now().Format(xtime.DateTime)),
	})
	if err != nil {
		log.Errorf("marshal request data failed: %v", err)
		return
	}

	request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(b))
	if err != nil {
		log.Errorf("create http request failed: %v", err)
		return
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Errorf("send http request failed: %v", err)
		return
	}
	defer response.Body.Close()

	if b, err = io.ReadAll(response.Body); err != nil {
		log.Errorf("read response body failed: %v", err)
		return
	}

	res := &greetRes{}
	resp := &duehttp.Resp{Data: res}

	if err = json.Unmarshal(b, resp); err != nil {
		log.Errorf("ummarshal response data failed: %v", err)
	}

	if resp.Code != 0 {
		log.Errorf("web server response failed, code: %d", resp.Code)
		return
	}

	log.Info(res.Message)
}
