package payload

type GemTotal struct {
	UserID uint   `json:"userId"`
	Total  uint64 `json:"total"`
}
