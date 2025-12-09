package person

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func NewController(s Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *Controller) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	person, err := c.service.GetByID(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "person not found"})
		return
	}

	ctx.JSON(http.StatusOK, person)
}

func (c *Controller) Create(ctx *gin.Context) {
	var p Person
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.service.Create(p)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, result)
}

func (c *Controller) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var p Person
	ctx.ShouldBindJSON(&p)

	result, err := c.service.Update(id, p)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *Controller) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Delete"})
}
