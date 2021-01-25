package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/carrepairrecord"
	"github.com/team07/app/ent/user"
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/repairing"
)

// CarRepairrecordController defines the struct for the carrepairrecord controller
type CarRepairrecordController struct {
	client *ent.Client
	router gin.IRouter
}

// CarRepairrecord defines the struct for the carrepairrecord
type CarRepairrecord struct {
	RepairingID		int
	UserID			int
	CarInspectionID	int
	Datetime		string
	Partrepair		string
	Price			int
	Techniciancomment string
}

// CreateCarRepairrecord handles POST requests for adding carrepairrecord entities
// @Summary Create carrepairrecord
// @Description Create carrepairrecord
// @ID create-carrepairrecord
// @Accept   json
// @Produce  json
// @Param carrepairrecord body CarRepairrecord true "CarRepairrecord entity"
// @Success 200 {object} ent.CarRepairrecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carrepairrecords [post]
func (ctl *CarRepairrecordController) CreateCarRepairrecord(c *gin.Context) {
	obj := CarRepairrecord{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "carrepairrecord binding failed",
		})
		return
	}

	r, err := ctl.client.Repairing.
		Query().
		Where(repairing.IDEQ(int(obj.RepairingID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "repairing not found",
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

	ci, err := ctl.client.CarInspection.
		Query().
		Where(carinspection.IDEQ(int(obj.CarInspectionID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "carinspection not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.Datetime)

	cr, err := ctl.client.CarRepairrecord.
		Create().
		SetCarinspection(ci).
		SetUser(u).
		SetKeeper(r).
		SetDatetime(times).
		SetPartrepair(obj.Partrepair).
		SetPrice(obj.Price).
		SetTechniciancomment(obj.Techniciancomment).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
			"status": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": cr,
	})
}

// GetCarRepairrecord handles GET requests to retrieve a carrepairrecord entity
// @Summary Get a carrepairrecord entity by ID
// @Description get carrepairrecord by ID
// @ID get-carrepairrecord
// @Produce  json
// @Param id path int true "CarRepairrecord ID"
// @Success 200 {object} ent.CarRepairrecord
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carrepairrecords/{id} [get]
func (ctl *CarRepairrecordController) GetCarRepairrecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ci, err := ctl.client.CarRepairrecord.
		Query().
		WithCarinspection().
		WithUser().
		WithKeeper().
		Where(carrepairrecord.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ci)
}

// ListCarRepairrecord handles request to get a list of carrepairrecord entities
// @Summary List carrepairrecord entities
// @Description list carrepairrecord entities
// @ID list-carrepairrecord
// @Produce json
// @Success 200 {array} ent.CarRepairrecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carrepairrecords [get]
func (ctl *CarRepairrecordController) ListCarRepairrecord(c *gin.Context) {
	cr, err := ctl.client.CarRepairrecord.
		Query().
		WithCarinspection().
		WithUser().
		WithKeeper().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cr)
}

// DeleteCarRepairrecord handles DELETE requests to delete a carrepairrecord entity
// @Summary Delete a carrepairrecord entity by ID
// @Description get carrepairrecord by ID
// @ID delete-carrepairrecord
// @Produce  json
// @Param id path int true "CarRepairrecord ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carrepairrecords/{id} [delete]
func (ctl *CarRepairrecordController) DeleteCarRepairrecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.CarRepairrecord.
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

// NewCarRepairrecord creates and registers handles for the carrepairrecord controller
func NewCarRepairrecord(router gin.IRouter, client *ent.Client) *CarRepairrecordController {
	cr := &CarRepairrecordController{
		client: client,
		router: router,
	}
	cr.register()
	return cr
}

// InitCarRepairrecordController register routes to the main engine
func (ctl *CarRepairrecordController) register() {
	cr := ctl.router.Group("/carrepairrecords")

	cr.GET("", ctl.ListCarRepairrecord)
	cr.POST("", ctl.CreateCarRepairrecord)
	// CRUD
	cr.GET(":id", ctl.GetCarRepairrecord)
	cr.DELETE(":id", ctl.DeleteCarRepairrecord)
}