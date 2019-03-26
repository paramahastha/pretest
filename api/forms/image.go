package forms

//ImageForm ...
type ImageForm struct {
	Name   string `form:"name" json:"name" binding:"required,max=64"`
	File   string `form:"file" json:"file" binding:"required,max=255"`
	Enable *bool  `form:"enable" json:"enable" binding:"exists"`
}
