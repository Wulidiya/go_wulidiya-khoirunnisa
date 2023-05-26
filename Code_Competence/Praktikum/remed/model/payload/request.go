package payload

type CreateItemRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Stock       uint   `json:"stock" form:"name" stock:"required"`
	Price       uint   `json:"price" form:"price" validate:"required"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"required"`
}

type CategoryRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}
