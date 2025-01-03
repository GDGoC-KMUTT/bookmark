package payload

type StrengthDataResponse struct {
	Data     []StrengthFieldData `json:"data"`
	Username string              `json:"username"`
}

type StrengthFieldData struct {
	FieldName string `json:"fieldName"`
	TotalGems int64  `json:"totalGems"`
}
