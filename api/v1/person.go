package v1

type PersonRequest struct {
	ID int64 `json:"id,string" binding:"id"`
}
