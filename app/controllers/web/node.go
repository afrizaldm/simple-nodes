package web

import (
	"net/http"
	"nodes_app/app/controllers"
	"nodes_app/database"
	"nodes_app/database/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Node struct {
	*controllers.Controller
}

type NodePost struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	ParentID uint   `json:"parent_id"`
}

func (w *Node) Index(c *gin.Context) {

	var data []models.Node
	result := database.DB.Find(&data)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":   data,
		"status": "success",
	})

}

func (w *Node) Create(c *gin.Context) {
	c.String(http.StatusOK, "ping")
}

func (w *Node) Store(c *gin.Context) {

	var data NodePost

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := database.DB.Create(&models.Node{
		Name:     data.Name,
		ParentID: data.ParentID,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data saved",
	})

}

func (w *Node) Show(c *gin.Context) {

	id := c.Param("id")

	var data models.Node
	result := database.DB.First(&data, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   data,
		"status": "success",
	})
}

func (w *Node) Edit(c *gin.Context) {
	c.String(http.StatusOK, "ping")
}

func (w *Node) Update(c *gin.Context) {

	id := c.Param("id")

	var dataPost NodePost

	if err := c.ShouldBindJSON(&dataPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var data models.Node
	result := database.DB.First(&data, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	result = database.DB.Model(&data).Updates(models.Node{
		Name:     dataPost.Name,
		ParentID: dataPost.ParentID,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Data can't be saved.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data saved.",
	})
}

func (w *Node) Delete(c *gin.Context) {

	// id, _ := strconv.ParseUint(c.Param("id"), 10, 0)
	// var product models.Node
	// result := database.DB.Delete(&product, uint(id))

	// if result.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"status":  "error",
	// 		"message": "Data can't be deleted.",
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"status":  "success",
	// 	"message": "Data deleted.",
	// })

	id, _ := strconv.ParseUint(c.Param("id"), 10, 0)

	// Menghapus node dengan `id = 1` atau `parent_id = 1`
	result := database.DB.Where("id = ? OR parent_id = ?", uint(id), uint(id)).Delete(&models.Node{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Data can't be deleted.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data deleted.",
	})
}
