package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/distance"
)

// DistanceController defines the struct for the distance controller
type DistanceController struct {
	client *ent.Client
	router gin.IRouter
}

// GetDistance handles GET requests to retrieve a distance entity
// @Summary Get a distance entity by ID
// @Description get distance by ID
// @ID get-distance
// @Produce  json
// @Param id path int true "Distance ID"
// @Success 200 {object} ent.Distance
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /distances/{id} [get]
func (ctl *DistanceController) GetDistance(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	dis, err := ctl.client.Distance.
		Query().		
		Where(distance.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, dis)
}

// ListDistance handles request to get a list of distance entities
// @Summary List distance entities
// @Description list distance entities
// @ID list-distance
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Distance
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /distances [get]
func (ctl *DistanceController) ListDistance(c *gin.Context) {
	
	dis, err := ctl.client.Distance.
		Query().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dis)
}

// NewDistanceController creates and registers handles for the distance controller
func NewDistanceController(router gin.IRouter, client *ent.Client) *DistanceController {
	distance := &DistanceController{
		client: client,
		router: router,
	}
	distance.register()
	return distance
}

// InitDistanceController registers routes to the main engine
func (ctl *DistanceController) register() {
distances := ctl.router.Group("/distances")

	distances.GET("", ctl.ListDistance)

	// CRUD
	distances.GET(":id", ctl.GetDistance)
}