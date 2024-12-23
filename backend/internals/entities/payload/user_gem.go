package payload

type GemTotal struct {
	UserID uint   `json:"user_id"`
	Total  uint64 `json:"total"`
}
