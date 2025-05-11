package utils

type Response struct {
	Msg       string      `json:"message"`
	Data      interface{} `json:"data"`
	SatusCode int         `json:"status_code"`
}
