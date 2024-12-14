package payload

type Profile struct {
	Id        *uint64 `json:"id"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	PhotoUrl  *string `json:"photoUrl"`
}
