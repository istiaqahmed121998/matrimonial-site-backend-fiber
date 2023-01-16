package utils

type ResponseData struct {
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
	Error   interface{} `json:"error"`
}

func Response(res ResponseData) ResponseData {
	return res
}
