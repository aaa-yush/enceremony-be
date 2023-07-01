package commons

type GenericErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type GenericSuccessResponse struct {
	Status string `json:"status"`
}
