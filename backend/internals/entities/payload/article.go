package payload

type Article struct {
	Id    *uint64 `json:"id"`
	Title *string `json:"title"`
	Href  *string `json:"href"`
}
