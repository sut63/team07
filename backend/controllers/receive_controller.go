package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/receive"
)

// ReceiveController defines the struct for the receive controller
type ReceiveController struct {
	client *ent.Client
	router gin.IRouter
}

// GetReceive handles GET requests to retrieve a receive entity
// @Summary Get a receive entity by ID
// @Description get receive by ID
// @ID get-receive
// @Produce  json
// @Param id path int true "Receive ID"
// @Success 200 {object} ent.Receive
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /receives/{id} [get]
func (ctl *ReceiveController) GetReceive(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	r, err := ctl.client.Receive.
		Query().
		Where(receive.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// ListReceive handles request to get a list of receive entities
// @Summary List receive entities
// @Description list receive entities
// @ID list-receive
// @Produce json
// @Success 200 {array} ent.Receive
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /receives [get]
func (ctl *ReceiveController) ListReceive(c *gin.Context) {
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

	r, err := ctl.client.Receive.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, r)
}

// DeleteReceive handles DELETE requests to delete a receive entity
// @Summary Delete a receive entity by ID
// @Description get receive by ID
// @ID delete-receive
// @Produce  json
// @Param id path int true "Receive ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /receives/{id} [delete]
func (ctl *ReceiveController) DeleteReceive(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Receive.
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

// NewReceiveController creates and registers handles for the receive controller
func NewReceiveController(router gin.IRouter, client *ent.Client) *ReceiveController {
	receives := &ReceiveController{
		client: client,
		router: router,
	}
	receives.register()
	return receives
}

// InitReceiveController registers routes to the main engine
func (ctl *ReceiveController) register() {
	receives := ctl.router.Group("/receives")

	receives.GET("", ctl.ListReceive)

	// CRUD
	receives.GET(":id", ctl.GetReceive)
	receives.DELETE(":id", ctl.DeleteReceive)
}
