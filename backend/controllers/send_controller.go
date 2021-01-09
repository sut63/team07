package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/send"
)

// SendController defines the struct for the send controller
type SendController struct {
	client *ent.Client
	router gin.IRouter
}

// GetSend handles GET requests to retrieve a send entity
// @Summary Get a send entity by ID
// @Description get send by ID
// @ID get-send
// @Produce  json
// @Param id path int true "Send ID"
// @Success 200 {object} ent.Send
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /sends/{id} [get]
func (ctl *SendController) GetSend(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.Send.
		Query().
		Where(send.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListSend handles request to get a list of send entities
// @Summary List send entities
// @Description list send entities
// @ID list-send
// @Produce json
// @Success 200 {array} ent.Send
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /sends [get]
func (ctl *SendController) ListSend(c *gin.Context) {
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

	s, err := ctl.client.Send.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, s)
}

// DeleteSend handles DELETE requests to delete a send entity
// @Summary Delete a send entity by ID
// @Description get send by ID
// @ID delete-send
// @Produce  json
// @Param id path int true "Send ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /sends/{id} [delete]
func (ctl *SendController) DeleteSend(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Send.
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

// NewSendController creates and registers handles for the send controller
func NewSendController(router gin.IRouter, client *ent.Client) *SendController {
	sends := &SendController{
		client: client,
		router: router,
	}
	sends.register()
	return sends
}

// InitSendController registers routes to the main engine
func (ctl *SendController) register() {
	sends := ctl.router.Group("/sends")

	sends.GET("", ctl.ListSend)

	// CRUD
	sends.GET(":id", ctl.GetSend)
	sends.DELETE(":id", ctl.DeleteSend)
}
