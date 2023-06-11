package dto

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ToDoCreateReq struct {
	Todo   string `json:"todo"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type ToDoCreateRes struct {
	ID     int64  `json:"id"`
	Todo   string `json:"todo"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type ToDoUpdateReq struct {
	ID     int64  `json:"id"`
	Todo   string `json:"todo"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
