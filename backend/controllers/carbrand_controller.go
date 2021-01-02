package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/carbrand"
)

// CarbrandController defines the struct for the carbrand controller
type CarbrandController struct {
	client *ent.Client
	router gin.IRouter
}

// GetCarbrand handles GET requests to retrieve a carbrand entity
// @Summary Get a carbrand entity by ID
// @Description get carbrand by ID
// @ID get-carbrand
// @Produce  json
// @Param id path int true "Carbrand ID"
// @Success 200 {object} ent.Carbrand
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carbrands/{id} [get]
func (ctl *CarbrandController) GetCarbrand(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	carbrand, err := ctl.client.Carbrand.
		Query().
		Where(carbrand.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, carbrand)
}

// ListCarbrand handles request to get a list of carbrand entities
// @Summary List carbrand entities
// @Description list carbrand entities
// @ID list-carbrand
// @Produce json
// @Success 200 {array} ent.Carbrand
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carbrands [get]
func (ctl *CarbrandController) ListCarbrand(c *gin.Context) {
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

	carbrands, err := ctl.client.Carbrand.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, carbrands)
}

// DeleteCarbrand handles DELETE requests to delete a carbrand entity
// @Summary Delete a carbrand entity by ID
// @Description get carbrand by ID
// @ID delete-carbrand
// @Produce  json
// @Param id path int true "Carbrand ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carbrands/{id} [delete]
func (ctl *CarbrandController) DeleteCarbrand(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Carbrand.
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

// NewCarbrandController creates and registers handles for the carbrand controller
func NewCarbrandController(router gin.IRouter, client *ent.Client) *CarbrandController {
	carbrands := &CarbrandController{
		client: client,
		router: router,
	}
	carbrands.register()
	return carbrands
}

// InitCarbrandController registers routes to the main engine
func (ctl *CarbrandController) register() {
	carbrands := ctl.router.Group("/carbrands")

	carbrands.GET("", ctl.ListCarbrand)

	// CRUD
	carbrands.GET(":id", ctl.GetCarbrand)
	carbrands.DELETE(":id", ctl.DeleteCarbrand)
}
