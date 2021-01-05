package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/urgent"
)

// UrgentController defines the struct for the urgent controller
type UrgentController struct {
	client *ent.Client
	router gin.IRouter
}

// GetUrgent handles GET requests to retrieve a urgent entity
// @Summary Get a urgent entity by ID
// @Description get urgent by ID
// @ID get-urgent
// @Produce  json
// @Param id path int true "Urgent ID"
// @Success 200 {object} ent.Urgent
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carservices/{id} [get]
func (ctl *UrgentController) GetUrgent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ug, err := ctl.client.Urgent.
		Query().		
		Where(urgent.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ug)
}

// ListUrgent handles request to get a list of urgent entities
// @Summary List urgent entities
// @Description list urgent entities
// @ID list-urgent
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Urgent
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /urgents [get]
func (ctl *UrgentController) ListUrgent(c *gin.Context) {
	
	ug, err := ctl.client.Urgent.
		Query().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, ug)
}

// NewUrgentController creates and registers handles for the urgent controller
func NewUrgentController(router gin.IRouter, client *ent.Client) *UrgentController {
	urgent := &UrgentController{
		client: client,
		router: router,
	}
	urgent.register()
	return urgent
}

// InitUrgentController registers routes to the main engine
func (ctl *UrgentController) register() {
	urgents := ctl.router.Group("/urgents")

	urgents.GET("", ctl.ListUrgent)

	// CRUD
	urgents.GET(":id", ctl.GetUrgent)
}