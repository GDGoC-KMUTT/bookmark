package payload

type ModuleStepResponse struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
	Check string `json:"check"`
}
