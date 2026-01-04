package backend

type ResponseData struct {
	Result  any     `json:"result"`
	Success bool    `json:"success"`
	Error   *string `json:"error"`
	// UnAuthorizedRequest bool    `json:"unAuthorizedRequest"`
}

type RespList struct {
	TotalCount int `json:"total_count"`
	Items      any `json:"items"`
}
