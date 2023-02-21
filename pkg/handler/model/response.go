package model

type JSONGeneralResponse struct {
	ServiceCode int         `json:"serverCode"`
	Msg         string      `json:"msg"`
	Data        interface{} `json:"data,omitempty"`
}
