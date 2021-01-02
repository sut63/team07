package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/insurance"
	"github.com/team07/app/ent/user"
	"github.com/team07/app/ent/carbrand"
	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/inspectionresult"
)

// AmbulanceController defines the struct for the ambulance controller
type AmbulanceController struct {
	client *ent.Client
	router gin.IRouter
}

// Ambulance defines the struct for the ambulance
type Ambulance struct {
	CarstatusID int
	CarbrandID int
	InsuranceID int
	UserID     int
	Registration  string
	Datetime  string
}

// CreateAmbulance handles POST requests for adding ambulance entities
// @Summary Create ambulance
// @Description Create ambulance
// @ID create-ambulance
// @Accept   json
// @Produce  json
// @Param ambulance body Ambulance true "Ambulance entity"
// @Success 200 {object} ent.Ambulance
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ambulances [post]
func (ctl *AmbulanceController) CreateAmbulance(c *gin.Context) {
	obj := Ambulance{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "ambulance binding failed",
		})
		return
	}

	ir, err := ctl.client.InspectionResult.
		Query().
		Where(inspectionresult.IDEQ(int(obj.CarstatusID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "carstatus not found",
		})
		return
	}

	carbrand, err := ctl.client.Carbrand.
		Query().
		Where(carbrand.IDEQ(int(obj.CarbrandID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "carbrand not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.UserID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}
	insurance, err := ctl.client.Insurance.
		Query().
		Where(insurance.IDEQ(int(obj.InsuranceID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "insurance not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.Datetime)

	am, err := ctl.client.Ambulance.
		Create().
		SetHasbrand(carbrand).
		SetHasuser(u).
		SetHasstatus(ir).
		SetRegisterat(times).
		SetHasinsurance(insurance).
		SetCarregistration(obj.Registration).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, am)
}

// GetAmbulance handles GET requests to retrieve a ambulance entity
// @Summary Get a ambulance entity by ID
// @Description get ambulance by ID
// @ID get-ambulance
// @Produce  json
// @Param id path int true "Ambulance ID"
// @Success 200 {object} ent.Ambulance
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ambulances/{id} [get]
func (ctl *AmbulanceController) GetAmbulance(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	am, err := ctl.client.Ambulance.
		Query().
		WithHasstatus().
		WithHasbrand().
		WithHasuser().
        WithHasinsurance().
		Where(ambulance.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, am)
}

// ListAmbulance handles request to get a list of ambulance entities
// @Summary List ambulance entities
// @Description list ambulance entities
// @ID list-ambulance
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Ambulance
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ambulances [get]
func (ctl *AmbulanceController) ListAmbulance(c *gin.Context) {
	
    ambulances, err := ctl.client.Ambulance.
		Query().
		WithHasstatus().
		WithHasbrand().
		WithHasuser().
        WithHasinsurance().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, ambulances)
}

// DeleteAmbulance handles DELETE requests to delete a Ambulance entity
// @Summary Delete a Ambulance entity by ID
// @Description get Ambulance by ID
// @ID delete-Ambulance
// @Produce  json
// @Param id path int true "Ambulance ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ambulances/{id} [delete]
func (ctl *AmbulanceController) DeleteAmbulance(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Ambulance.
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

// NewAmbulanceController creates and registers handles for the ambulance controller
func NewAmbulanceController(router gin.IRouter, client *ent.Client) *AmbulanceController {
	ambulance := &AmbulanceController{
		client: client,
		router: router,
	}
	ambulance.register()
	return ambulance
}

// InitAmbulanceController registers routes to the main engine
func (ctl *AmbulanceController) register() {
	ambulances := ctl.router.Group("/ambulances")

	ambulances.GET("", ctl.ListAmbulance)

	// CRUD
	ambulances.GET(":id", ctl.GetAmbulance)
	ambulances.POST("", ctl.CreateAmbulance)
	ambulances.DELETE(":id", ctl.DeleteAmbulance)
}
