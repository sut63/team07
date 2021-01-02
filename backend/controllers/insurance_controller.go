package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/insurance"
)

// InsuranceController defines the struct for the insurance controller
type InsuranceController struct {
	client *ent.Client
	router gin.IRouter
}

// GetInsurance handles GET requests to retrieve a insurance entity
// @Summary Get a insurance entity by ID
// @Description get insurance by ID
// @ID get-insurance
// @Produce  json
// @Param id path int true "Insurance ID"
// @Success 200 {object} ent.Insurance
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /insurances/{id} [get]
func (ctl *InsuranceController) GetInsurance(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	insurance, err := ctl.client.Insurance.
		Query().
		Where(insurance.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, insurance)
}

// ListInsurance handles request to get a list of insurance entities
// @Summary List insurance entities
// @Description list insurance entities
// @ID list-insurance
// @Produce json
// @Success 200 {array} ent.Insurance
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /insurances [get]
func (ctl *InsuranceController) ListInsurance(c *gin.Context) {
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

	insurances, err := ctl.client.Insurance.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, insurances)
}

// DeleteInsurance handles DELETE requests to delete a insurance entity
// @Summary Delete a insurance entity by ID
// @Description get insurance by ID
// @ID delete-insurance
// @Produce  json
// @Param id path int true "Insurance ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /insurances/{id} [delete]
func (ctl *InsuranceController) DeleteInsurance(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Insurance.
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

// NewInsuranceController creates and registers handles for the insurance controller
func NewInsuranceController(router gin.IRouter, client *ent.Client) *InsuranceController {
	insurances := &InsuranceController{
		client: client,
		router: router,
	}
	insurances.register()
	return insurances
}

// InitInsuranceController registers routes to the main engine
func (ctl *InsuranceController) register() {
	insurances := ctl.router.Group("/insurances")

	insurances.GET("", ctl.ListInsurance)

	// CRUD
	insurances.GET(":id", ctl.GetInsurance)
	insurances.DELETE(":id", ctl.DeleteInsurance)
}
