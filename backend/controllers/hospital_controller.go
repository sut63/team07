package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/hospital"
)

// HospitalController defines the struct for the hospital controller
type HospitalController struct {
	client *ent.Client
	router gin.IRouter
}

// GetHospital handles GET requests to retrieve a hospital entity
// @Summary Get a hospital entity by ID
// @Description get hospital by ID
// @ID get-hospital
// @Produce  json
// @Param id path int true "Hospital ID"
// @Success 200 {object} ent.Hospital
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /hospitals/{id} [get]
func (ctl *HospitalController) GetHospital(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.Hospital.
		Query().
		Where(hospital.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListHospital handles request to get a list of hospital entities
// @Summary List hospital entities
// @Description list hospital entities
// @ID list-hospital
// @Produce json
// @Success 200 {array} ent.Hospital
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /hospitals [get]
func (ctl *HospitalController) ListHospital(c *gin.Context) {
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

	s, err := ctl.client.Hospital.
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

// DeleteHospital handles DELETE requests to delete a hospital entity
// @Summary Delete a hospital entity by ID
// @Description get hospital by ID
// @ID delete-hospital
// @Produce  json
// @Param id path int true "Hospital ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /hospitals/{id} [delete]
func (ctl *HospitalController) DeleteHospital(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Hospital.
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

// NewHospitalController creates and registers handles for the hospital controller
func NewHospitalController(router gin.IRouter, client *ent.Client) *HospitalController {
	hospitals := &HospitalController{
		client: client,
		router: router,
	}
	hospitals.register()
	return hospitals
}

// InitHospitalController registers routes to the main engine
func (ctl *HospitalController) register() {
	hospitals := ctl.router.Group("/hospitals")

	hospitals.GET("", ctl.ListHospital)

	// CRUD
	hospitals.GET(":id", ctl.GetHospital)
	hospitals.DELETE(":id", ctl.DeleteHospital)
}