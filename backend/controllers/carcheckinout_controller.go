package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/user"
	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/purpose"
	"github.com/team07/app/ent/carcheckinout"
)

// CarCheckInOutController defines the struct for the carcheckinout controller
type CarCheckInOutController struct {
	client *ent.Client
	router gin.IRouter
}

// CarCheckInOut defines the struct for the carcheckinout
type CarCheckInOut struct {
	Ambulance   int
	Name    	int
	Purpose    	int
	Note   		string
	Checkin     string
	Checkout    string
}

// CreateCarCheckInOut handles POST requests for adding carcheckinout entities
// @Summary Create carcheckinout
// @Description Create carcheckinout
// @ID create-carcheckinout
// @Accept   json
// @Produce  json
// @Param carcheckinout body CarCheckInOut true "CarCheckInOut entity"
// @Success 200 {object} CarCheckInOut
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carcheckinouts [post]
func (ctl *CarCheckInOutController) CreateCarCheckInOut(c *gin.Context) {
	obj := CarCheckInOut{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "carcheckinout binding failed",
		})
		return
	}

	a, err := ctl.client.Ambulance.
		Query().
		Where(ambulance.IDEQ(int(obj.Ambulance))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ambulance not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.Name))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	p, err := ctl.client.Purpose.
		Query().
		Where(purpose.IDEQ(int(obj.Purpose))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "purpose not found",
		})
		return
	}

	timein, err := time.Parse(time.RFC3339, obj.Checkin)
	timeout, err := time.Parse(time.RFC3339, obj.Checkout)

	cio, err := ctl.client.CarCheckInOut.
		Create().
		SetAmbulance(a).
		SetName(u).
		SetPurpose(p).
		SetNote(obj.Note).
		SetCheckIn(timein).
		SetCheckOut(timeout).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, cio)
}

// GetCarCheckInOut handles GET requests to retrieve a carcheckinout entity
// @Summary Get a carcheckinout entity by ID
// @Description get carcheckinout by ID
// @ID get-carcheckinout
// @Produce  json
// @Param id path int true "CarCheckInOut ID"
// @Success 200 {object} ent.CarCheckInOut
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carcheckinouts/{id} [get]
func (ctl *CarCheckInOutController) GetCarCheckInOut(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cio, err := ctl.client.CarCheckInOut.
		Query().
		Where(carcheckinout.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cio)
}

// ListCarCheckInOut handles request to get a list of carcheckinout entities
// @Summary List carcheckinout entities
// @Description list carcheckinout entities
// @ID list-carcheckinout
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.CarCheckInOut
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carcheckinouts [get]
func (ctl *CarCheckInOutController) ListCarCheckInOut(c *gin.Context) {
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

	carcheckinouts, err := ctl.client.CarCheckInOut.
		Query().
		WithAmbulance().
		WithName().
		WithPurpose().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, carcheckinouts)
}

// DeleteCarCheckInOut handles DELETE requests to delete a carcheckinout entity
// @Summary Delete a carcheckinout entity by ID
// @Description get carcheckinout by ID
// @ID delete-carcheckinout
// @Produce  json
// @Param id path int true "CarCheckInOut ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carcheckinouts/{id} [delete]
func (ctl *CarCheckInOutController) DeleteCarCheckInOut(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.CarCheckInOut.
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

// UpdateCarCheckInOut handles PUT requests to update a carcheckinout entity
// @Summary Update a carcheckinout entity by ID
// @Description update carcheckinout by ID
// @ID update-carcheckinout
// @Accept   json
// @Produce  json
// @Param id path int true "CarCheckInOut ID"
// @Param carcheckinout body ent.CarCheckInOut true "CarCheckInOut entity"
// @Success 200 {object} ent.CarCheckInOut
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carcheckinouts/{id} [put]
func (ctl *CarCheckInOutController) UpdateCarCheckInOut(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.CarCheckInOut{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "carcheckinout binding failed",
		})
		return
	}
	
	obj.ID = int(id)
	cio, err := ctl.client.CarCheckInOut.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, cio)
}

// NewCarCheckInOutController creates and registers handles for the carcheckinout controller
func NewCarCheckInOutController(router gin.IRouter, client *ent.Client) *CarCheckInOutController {
	uc := &CarCheckInOutController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitCarCheckInOutController registers routes to the main engine
func (ctl *CarCheckInOutController) register() {
	carcheckinouts := ctl.router.Group("/carcheckinouts")

	carcheckinouts.GET("", ctl.ListCarCheckInOut)

	// CRUD
	carcheckinouts.POST("", ctl.CreateCarCheckInOut)
	carcheckinouts.GET(":id", ctl.GetCarCheckInOut)
	carcheckinouts.PUT(":id", ctl.UpdateCarCheckInOut)
	carcheckinouts.DELETE(":id", ctl.DeleteCarCheckInOut)
}
