package server

type errorResponse struct {
	Error string `json:"error"`
}

type uploadFileResponse struct {
	Url string `json:"url"`
}
