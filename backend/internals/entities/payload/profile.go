package payload

type Profile struct {
	Id        *uint64 `json:"id"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Email     *string `json:"email"`
	PhotoUrl  *string `json:"photoUrl"`
}
