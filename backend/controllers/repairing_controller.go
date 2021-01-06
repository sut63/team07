package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/repairing"
)

// RepairingController defines the struct for the repair controller
type RepairingController struct {
	client *ent.Client
	router gin.IRouter
}

// GetRepairing handles GET requests to retrieve a repairing entity
// @Summary Get a repairing entity by ID
// @Description get repairing by ID
// @ID get-repairing
// @Produce  json
// @Param id path int true "Repairing ID"
// @Success 200 {object} ent.Repairing
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /repairings/{id} [get]
func (ctl *RepairingController) GetRepairing(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	r, err := ctl.client.Repairing.
		Query().
		Where(repairing.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// ListRepairing handles request to get a list of repairing entities
// @Summary List Repairing entities
// @Description list repairing entities
// @ID list-repairing
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Repairing
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /repairings [get]
func (ctl *RepairingController) ListRepairing(c *gin.Context) {

	r, err := ctl.client.Repairing.
		Query().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, r)
}

// NewRepairingController creates and reginters hanfles for te repairing controller
func NewRepairingController(router gin.IRouter, client *ent.Client) *RepairingController {
	repairing := &RepairingController{
		client: client,
		router: router,
	}
	repairing.register()
	return repairing
}

// InitRepairingController register routes to the main engine
func (ctl *RepairingController) register() {
	repairings := ctl.router.Group("/repairings")

	repairings.GET("", ctl.ListRepairing)

	// CRUD
	repairings.GET(":id", ctl.GetRepairing)
}