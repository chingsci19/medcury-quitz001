package base

type Response struct {
	StatusCode string      `json:"status_code" `
	StatusDesc string      `json:"status_description" `
	Data       interface{} `json:"data" `
	FileName   *string     `json:"file_name,omitempty" `
	HttpStatus int         `json:"-" `
}

type AppError struct {
	StatusCode string `json:"status_code" `
	Message    string `json:"message" `
	HttpStatus int    `json:"-" `
}
