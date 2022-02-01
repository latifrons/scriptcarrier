package rpc

type GeneralResponse struct {
	Code int         `json:"code" example:"0"` // Code is 0 for normal cases and positive for errors.
	Msg  string      `json:"msg"`              // Msg is "" for normal cases and message for errors.
	Data interface{} `json:"data,omitempty"`   // Optional
}

type PagingResponse struct {
	GeneralResponse
	List  interface{} `json:"list,omitempty"` // List is always the result list
	Size  int         `json:"size"`           // Size is the result count in this response1
	Total int64       `json:"total"`          // Total is the total result in database
	Page  int         `json:"page"`           // Page is the given params in the request starts from 1
}
