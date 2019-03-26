package controllers

import (
	"strconv"

	"github.com/paramahastha/pretest/api/forms"
	"github.com/paramahastha/pretest/api/models"

	"github.com/gin-gonic/gin"
)

//ProductController ...
type ProductController struct{}

var productModel = new(models.ProductModel)

//Create ...
func (ctrl ProductController) Create(c *gin.Context) {
	var productForm forms.ProductForm

	if c.BindJSON(&productForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": productForm})
		c.Abort()
		return
	}

	err := productModel.Create(productForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Product could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Product created"})
}

//All ...
func (ctrl ProductController) All(c *gin.Context) {
	data, err := productModel.All()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the products", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": data})
}

//One ...
func (ctrl ProductController) One(c *gin.Context) {
	id := c.Param("id")

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		data, err := productModel.One(id)

		if err != nil {
			c.JSON(404, gin.H{"Message": "Product not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Update ...
func (ctrl ProductController) Update(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		var productForm forms.ProductForm

		if c.BindJSON(&productForm) != nil {
			c.JSON(406, gin.H{"message": "Invalid parameters", "form": productForm})
			c.Abort()
			return
		}

		err := productModel.Update(id, productForm)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Product could not be updated", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Product updated"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
	}
}

//Delete ...
func (ctrl ProductController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		err := productModel.Delete(id)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Product could not be deleted", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Product deleted"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
