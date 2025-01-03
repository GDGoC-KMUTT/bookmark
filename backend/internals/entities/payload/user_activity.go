package payload

type UserActivityParam struct {
    UserId   uint64 `json:"userId"`
    StepId uint64 `json:"stepId"`
}
