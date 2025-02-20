package handler

import (
	"context"
	"net/http"
	"strconv"

	"go-brain/internal/model"
	"go-brain/internal/service"

	"github.com/gin-gonic/gin"
)

type ThingHandler struct {
	ThingService *service.ThingService
}

func NewThingHandler(thingService *service.ThingService) *ThingHandler {
	return &ThingHandler{ThingService: thingService}
}

func (h *ThingHandler) Create(c *gin.Context) {
	var thing model.Thing
	if err := c.ShouldBindJSON(&thing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.ThingService.Create(context.Background(), &thing)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, thing)
}

func (h *ThingHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	thing, err := h.ThingService.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "thing not found"})
		return
	}
	c.JSON(http.StatusOK, thing)
}

func (h *ThingHandler) Update(c *gin.Context) {
	var thing model.Thing
	if err := c.ShouldBindJSON(&thing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	thing.ID = id
	err = h.ThingService.Update(context.Background(), &thing)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, thing)
}

func (h *ThingHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	err = h.ThingService.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "thing deleted"})
}
