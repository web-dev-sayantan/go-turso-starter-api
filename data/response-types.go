package data

type SuccessResType[T any] struct {
	Success bool         `json:"success"`
	Data    map[string]T `json:"data"`
}

type FailureResType struct {
	Success bool   `json:"success"`
	Error   string `json:"data"`
}
