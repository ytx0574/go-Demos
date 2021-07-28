package Const

import "encoding/json"

type HTTPResponseCode int
type HTTPResponseMessage string

const (
	Code string = "code"
	Message string = "message"
	Data string = "data"
)

const (
	KSuccessCode HTTPResponseCode = 0
	KUserExsit HTTPResponseCode = 100
	KUserNotExsit HTTPResponseCode = 101

	KBookNotExsit HTTPResponseCode = 1001
	KBookStockShortage HTTPResponseCode = 1002

)

var HTTPMessage = map[HTTPResponseCode]HTTPResponseMessage {
	KUserExsit : "用户已存在",
	KUserNotExsit : "用户不存在",
	KBookNotExsit : "图书不存在",
	KBookStockShortage : "库存不足",
}

type ResponseMessage struct {
	Code HTTPResponseCode `json:"code"`
	Message HTTPResponseMessage `json:"message"`
	Info interface{} `json:"info"`
}

func NewResponseMessage(code HTTPResponseCode, info interface{}) ResponseMessage {
	return ResponseMessage{
		Code: code,
		Message: HTTPMessage[code],
		Info: info,
	}
}

func NewResponseCustomSuccessMessage(msg HTTPResponseMessage, info interface{}) ResponseMessage {
	return ResponseMessage{
		Code: KSuccessCode,
		Message: msg,
		Info: info,
	}
}

func NewJSONBytesResponseMessage(message ResponseMessage) []byte {

	bytes, err := json.Marshal(message)

	if err != nil {
		panic(err)
	}

	return bytes
}



