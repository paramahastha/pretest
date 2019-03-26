package controllers

import (
	"strconv"

	"github.com/paramahastha/pretest/api/forms"
	"github.com/paramahastha/pretest/api/models"

	"github.com/gin-gonic/gin"
)

//ImageController ...
type ImageController struct{}

var imageModel = new(models.ImageModel)

//Create ...
func (ctrl ImageController) Create(c *gin.Context) {
	var imageForm forms.ImageForm

	if c.BindJSON(&imageForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": imageForm})
		c.Abort()
		return
	}

	err := imageModel.Create(imageForm)

	if err != nil {
		c.JSON(406, gin.H{"message": "Image could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Image created"})
}

//All ...
func (ctrl ImageController) All(c *gin.Context) {
	data, err := imageModel.All()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the images", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": data})
}

//One ...
func (ctrl ImageController) One(c *gin.Context) {
	id := c.Param("id")

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		data, err := imageModel.One(id)

		if err != nil {
			c.JSON(404, gin.H{"Message": "Image not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Update ...
func (ctrl ImageController) Update(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		var imageForm forms.ImageForm

		if c.BindJSON(&imageForm) != nil {
			c.JSON(406, gin.H{"message": "Invalid parameters", "form": imageForm})
			c.Abort()
			return
		}

		err := imageModel.Update(id, imageForm)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Image could not be updated", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Image updated"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
	}
}

//Delete ...
func (ctrl ImageController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		err := imageModel.Delete(id)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Image could not be deleted", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Image deleted"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
