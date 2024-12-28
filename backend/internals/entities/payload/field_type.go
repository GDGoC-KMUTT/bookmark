package payload

type FieldType struct {
	Id       *uint64 `json:"id"`
	Name     *string `json:"name"`
	ImageUrl *string `json:"imageUrl"`
}

type FieldIdParam struct {
	FieldId uint `params:"fieldId"`
}
