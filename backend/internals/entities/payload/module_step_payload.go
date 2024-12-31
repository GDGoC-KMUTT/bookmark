package payload

type ModuleStep struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
	Check bool   `json:"check"`
}
