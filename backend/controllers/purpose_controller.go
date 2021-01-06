package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/purpose"
)

// PurposeController defines the struct for the purpose controller
type PurposeController struct {
	client *ent.Client
	router gin.IRouter
}

// GetPurpose handles GET requests to retrieve a purpose entity
// @Summary Get a purpose entity by ID
// @Description get purpose by ID
// @ID get-purpose
// @Produce  json
// @Param id path int true "Purpose ID"
// @Success 200 {object} ent.Purpose
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /purposes/{id} [get]
func (ctl *PurposeController) GetPurpose(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	purpose, err := ctl.client.Purpose.
		Query().
		Where(purpose.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, purpose)
}

// ListPurpose handles request to get a list of purpose entities
// @Summary List purpose entities
// @Description list purpose entities
// @ID list-purpose
// @Produce json
// @Success 200 {array} ent.Purpose
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /purposes [get]
func (ctl *PurposeController) ListPurpose(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	purposes, err := ctl.client.Purpose.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, purposes)
}

// DeletePurpose handles DELETE requests to delete a purpose entity
// @Summary Delete a purpose entity by ID
// @Description get purpose by ID
// @ID delete-purpose
// @Produce  json
// @Param id path int true "Purpose ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /purposes/{id} [delete]
func (ctl *PurposeController) DeletePurpose(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Purpose.
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

// NewPurposeController creates and registers handles for the purpose controller
func NewPurposeController(router gin.IRouter, client *ent.Client) *PurposeController {
	purposes := &PurposeController{
		client: client,
		router: router,
	}
	purposes.register()
	return purposes
}

// InitPurposeController registers routes to the main engine
func (ctl *PurposeController) register() {
	purposes := ctl.router.Group("/purposes")

	purposes.GET("", ctl.ListPurpose)

	// CRUD
	purposes.GET(":id", ctl.GetPurpose)
	purposes.DELETE(":id", ctl.DeletePurpose)
}
