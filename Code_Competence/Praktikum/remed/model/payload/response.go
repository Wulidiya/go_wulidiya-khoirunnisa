package payload

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name" form:"name" validate:"required"`
}
