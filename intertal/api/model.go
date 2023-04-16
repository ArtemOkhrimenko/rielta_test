package api

type CreateRequest struct {
	Number int64 `json:"number" binding:"gte=0"`
}
