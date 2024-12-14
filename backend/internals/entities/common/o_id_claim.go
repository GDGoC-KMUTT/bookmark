package common

type OIdClaims struct {
	Id        *string `json:"sub"`
	FirstName *string `json:"name"`
	Lastname  *string `json:"family_name"`
	Picture   *string `json:"picture"`
	Email     *string `json:"email"`
}
