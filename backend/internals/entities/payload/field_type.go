package payload

type FieldType struct {
	Id       *uint64 `json:"id"`
	Name     *uint64 `json:"name"`
	ImageUrl *string `json:"image_url"`
}

type FieldIdParam struct {
	FieldId *uint `params:"field_id"`
}
