package controllers

import (
	"fmt"
	"strconv"

	"github.com/paramahastha/pretest/api/forms"
	"github.com/paramahastha/pretest/api/models"

	"github.com/gin-gonic/gin"
)

//CategoryController ...
type CategoryController struct{}

var categoryModel = new(models.CategoryModel)

//Create ...
func (ctrl CategoryController) Create(c *gin.Context) {
	var categoryForm forms.CategoryForm

	if c.BindJSON(&categoryForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": categoryForm})
		c.Abort()
		return
	}

	err := categoryModel.Create(categoryForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Category could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Category created"})
}

//All ...
func (ctrl CategoryController) All(c *gin.Context) {
	data, err := categoryModel.All()
	
	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the categories", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": data})
}

//One ...
func (ctrl CategoryController) One(c *gin.Context) {
	id := c.Param("id")

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		data, err := categoryModel.One(id)
		fmt.Printf("%v %s", data, err)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Category not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Update ...
func (ctrl CategoryController) Update(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		var categoryForm forms.CategoryForm

		if c.BindJSON(&categoryForm) != nil {
			c.JSON(406, gin.H{"message": "Invalid parameters", "form": categoryForm})
			c.Abort()
			return
		}

		err := categoryModel.Update(id, categoryForm)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Category could not be updated", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Category updated"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
	}
}

//Delete ...
func (ctrl CategoryController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		err := categoryModel.Delete(id)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Category could not be deleted", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Category deleted"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
