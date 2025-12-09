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
	ctx.JSON(http.StatusOK, c.service.GetAll())
}

func (c *Controller) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	person := c.service.GetByID(id)

	if person == nil {
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

	result := c.service.Create(p)
	ctx.JSON(http.StatusCreated, result)
}

func (c *Controller) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var p Person
	ctx.ShouldBindJSON(&p)

	updated := c.service.Update(id, p)
	if updated == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (c *Controller) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	ok := c.service.Delete(id)

	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Person not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Delete"})
}
