package common

type OIdClaims struct {
	Id        *string `json:"sub"`
	FirstName *string `json:"firstName"`
	Lastname  *string `json:"lastName"`
	Picture   *string `json:"picture"`
	Email     *string `json:"email"`
}
