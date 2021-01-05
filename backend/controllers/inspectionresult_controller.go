package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/inspectionresult"
)

// InspectionResultController defines the struct for the inspectionresult controller
type InspectionResultController struct {
	client *ent.Client
	router gin.IRouter
}

// GetInspectionResult handles GET requests to retrieve a inspectionresult entity
// @Summary Get a inspectionresult entity by ID
// @Description get inspectionresult by ID
// @ID get-inspectionresult
// @Produce  json
// @Param id path int true "InspectionResult ID"
// @Success 200 {object} ent.InspectionResult
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /inspectionresults/{id} [get]
func (ctl *InspectionResultController) GetInspectionResult(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ir, err := ctl.client.InspectionResult.
		Query().
		WithJobposition().
		Where(inspectionresult.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ir)
}

// ListInspectionResult handles request to get a list of inspectionresult entities
// @Summary List inspectionresult entities
// @Description list inspectionresult entities
// @ID list-inspectionresult
// @Produce json
// @Success 200 {array} ent.InspectionResult
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /inspectionresults [get]
func (ctl *InspectionResultController) ListInspectionResult(c *gin.Context) {
	iresults, err := ctl.client.InspectionResult.
		Query().
		WithJobposition().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, iresults)
}

// DeleteInspectionResult handles DELETE requests to delete a inspectionresult entity
// @Summary Delete a inspectionresult entity by ID
// @Description get inspectionresult by ID
// @ID delete-inspectionresult
// @Produce  json
// @Param id path int true "InspectionResult ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /inspectionresults/{id} [delete]
func (ctl *InspectionResultController) DeleteInspectionResult(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.InspectionResult.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewInspectionResultController creates and registers handles for the user controller
func NewInspectionResultController(router gin.IRouter, client *ent.Client) *InspectionResultController {
	irc := &InspectionResultController{
		client: client,
		router: router,
	}
	irc.register()
	return irc
}

// InitInspectionResultController registers routes to the main engine
func (ctl *InspectionResultController) register() {
	iresults := ctl.router.Group("/inspectionresults")

	iresults.GET("", ctl.ListInspectionResult)

	// CRUD
	iresults.GET(":id", ctl.GetInspectionResult)
	iresults.DELETE(":id", ctl.DeleteInspectionResult)
}
