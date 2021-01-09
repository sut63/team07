package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/carservice"
	"github.com/team07/app/ent/user"
	"github.com/team07/app/ent/urgent"
	"github.com/team07/app/ent/distance"
)

// CarserviceController defines the struct for the carservice controller
type CarserviceController struct {
	client *ent.Client
	router gin.IRouter
}

// Carservice defines the struct for the carservice
type Carservice struct {
	DistanceID int
	UrgentID   int 
	UserID     int
	Customer  string
	Location  string
	Information  string 
	Datetime  string
}

// CreateCarservice handles POST requests for adding carservice entities
// @Summary Create carservice
// @Description Create carservice
// @ID create-carservice
// @Accept   json
// @Produce  json
// @Param carservice body Carservice true "Carservice entity"
// @Success 200 {object} ent.Carservice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carservices [post]
func (ctl *CarserviceController) CreateCarservice(c *gin.Context) {
	obj := Carservice{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "carservice binding failed",
		})
		return
	}

	di, err := ctl.client.Distance.
		Query().
		Where(distance.IDEQ(int(obj.DistanceID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "distance not found",
		})
		return
	}

	ug, err := ctl.client.Urgent.
		Query().
		Where(urgent.IDEQ(int(obj.UrgentID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "urgent not found",
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

	times, err := time.Parse(time.RFC3339, obj.Datetime)

	cars, err := ctl.client.Carservice.
		Create().
		SetUserid(u).
		SetUrgentid(ug).
		SetDisid(di).
		SetDatetime(times).
		SetCustomer(obj.Customer).
		SetLocation(obj.Location).
		SetInformation(obj.Information).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, cars)
}

// GetCarservice handles GET requests to retrieve a carservice entity
// @Summary Get a carservice entity by ID
// @Description get carservice by ID
// @ID get-carservice
// @Produce  json
// @Param id path int true "Carservice ID"
// @Success 200 {object} ent.Carservice
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carservices/{id} [get]
func (ctl *CarserviceController) GetCarservice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cars, err := ctl.client.Carservice.
		Query().		
		WithUserid().
		WithUrgentid().
		WithDisid().
		Where(carservice.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cars)
}

// ListCarservice handles request to get a list of carservice entities
// @Summary List carservice entities
// @Description list carservice entities
// @ID list-carservice
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Carservice
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carservices [get]
func (ctl *CarserviceController) ListCarservice(c *gin.Context) {
	
	cars, err := ctl.client.Carservice.
		Query().
		WithUserid().
		WithUrgentid().
		WithDisid().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cars)
}

// DeleteCarservice handles DELETE requests to delete a Carservice entity
// @Summary Delete a Carservice entity by ID
// @Description get Carservice by ID
// @ID delete-Carservice
// @Produce  json
// @Param id path int true "Carservice ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carservices/{id} [delete]
func (ctl *CarserviceController) DeleteCarservice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Carservice.
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

// NewCarserviceController creates and registers handles for the carservice controller
func NewCarserviceController(router gin.IRouter, client *ent.Client) *CarserviceController {
	carservice := &CarserviceController{
		client: client,
		router: router,
	}
	carservice.register()
	return carservice
}

// InitCarserviceController registers routes to the main engine
func (ctl *CarserviceController) register() {
	carservices := ctl.router.Group("/carservices")

	carservices.GET("", ctl.ListCarservice)

	// CRUD
	carservices.GET(":id", ctl.GetCarservice)
	carservices.POST("", ctl.CreateCarservice)
	carservices.DELETE(":id", ctl.DeleteCarservice)
}