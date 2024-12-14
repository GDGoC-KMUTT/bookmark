package payload

type OauthCallback struct {
	Code *string `json:"code" validate:"required"`
}
