package ws

import (
  "reflect"
  "github.com/gorilla/websocket"
  . "server/common"
  "server/common/utils"
)

type heartbeatPingReq struct {
  ClientTimestamp int64 `json:"clientTimestamp"`
}

func init()  {
  regHandleInfo("HeartbeatPing",
    &wsHandleInfo{reflect.TypeOf(heartbeatPingReq{}), "HeartbeatPong"})
}

func (req *heartbeatPingReq) handle(conn *websocket.Conn, resp *wsResp) error {
  resp.Ret = Constants.RetCode.Ok
  data := struct {
    ServerTimestamp int64 `json:"serverTimestamp"`
  }{utils.UnixtimeMilli()}
  resp.Data = data
  return nil
}
