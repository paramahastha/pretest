package forms

//CategoryForm ...
type CategoryForm struct {
	Name   string `form:"name" json:"name" binding:"required,max=64"`
	Enable *bool  `form:"enable" json:"enable" binding:"exists"`
}
