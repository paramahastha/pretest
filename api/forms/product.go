package forms

//ProductForm ...
type ProductForm struct {
	Name        string `form:"name" json:"name" binding:"required,max=64"`
	Description string `form:"description" json:"description" binding:"required,max=255"`
	Enable      *bool  `form:"enable" json:"enable" binding:"exists"`
}
