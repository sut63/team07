package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/ambulance"
	"github.com/team07/app/ent/carinspection"
	"github.com/team07/app/ent/inspectionresult"
	"github.com/team07/app/ent/user"
)

// CarInspectionController defines the struct for the carinspection controller
type CarInspectionController struct {
	client *ent.Client
	router gin.IRouter
}

// CarInspection defines the struct for the carinspection
type CarInspection struct {
	WheelCenter        float64
	BlackSmoke         float64
	SoundLevel         float64
	InspectionResultID int
	UserID             int
	AmbulanceID        int
	Datetime           string
	Note               string
}

// CreateCarInspection handles POST requests for adding carinspection entities
// @Summary Create carinspection
// @Description Create carinspection
// @ID create-carinspection
// @Accept   json
// @Produce  json
// @Param carinspection body CarInspection true "CarInspection entity"
// @Success 200 {object} ent.CarInspection
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carinspections [post]
func (ctl *CarInspectionController) CreateCarInspection(c *gin.Context) {
	obj := CarInspection{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "carinspection binding failed",
		})
		return
	}

	ir, err := ctl.client.InspectionResult.
		Query().
		Where(inspectionresult.IDEQ(int(obj.InspectionResultID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "inspectionresult not found",
		})
		return
	}

	a, err := ctl.client.Ambulance.
		Query().
		Where(ambulance.IDEQ(int(obj.AmbulanceID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ambulance not found",
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

	ci, err := ctl.client.CarInspection.
		Create().
		SetAmbulance(a).
		SetUser(u).
		SetInspectionresult(ir).
		SetDatetime(times).
		SetWheelCenter(obj.WheelCenter).
		SetBlacksmoke(obj.BlackSmoke).
		SetSoundLevel(obj.SoundLevel).
		SetNote(obj.Note).
		Save(context.Background())
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   ci,
	})
}

// GetCarInspection handles GET requests to retrieve a carinspection entity
// @Summary Get a carinspection entity by ID
// @Description get carinspection by ID
// @ID get-carinspection
// @Produce  json
// @Param id path int true "CarInspection ID"
// @Success 200 {object} ent.CarInspection
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carinspections/{id} [get]
func (ctl *CarInspectionController) GetCarInspection(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ci, err := ctl.client.CarInspection.
		Query().
		WithAmbulance().
		WithInspectionresult().
		WithUser().
		Where(carinspection.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ci)
}

// GetCarInspectionBySearch handles GET requests to retrieve a carinspection entity
// @Summary Get a carinspection entity by Search
// @Description get carinspection by Search
// @ID get-carinspection-by-search
// @Produce  json
// @Param ambulance query string false "Ambulance Search"
// @Param result query int false "Result Search"
// @Param user query string false "User Search"
// @Success 200 {object} ent.CarInspection
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /searchcarinspections [get]
func (ctl *CarInspectionController) GetCarInspectionBySearch(c *gin.Context) {
	asearch := c.Query("ambulance")
	usearch := c.Query("user")
	irsearch, err := strconv.ParseInt(c.Query("result"), 10, 64)

	irstring := ""
	ir, err := ctl.client.InspectionResult.
		Query().
		Where(inspectionresult.IDEQ(int(irsearch))).
		Only(context.Background())

	if ir != nil {
		irstring = ir.ResultName
	}

	ci, err := ctl.client.CarInspection.
		Query().
		WithAmbulance().
		WithInspectionresult().
		WithUser().
		Where(carinspection.HasUserWith(user.NameContains(usearch))).
		Where(carinspection.HasAmbulanceWith(ambulance.CarregistrationContains(asearch))).
		Where(carinspection.HasInspectionresultWith(inspectionresult.ResultNameContains(irstring))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	if asearch == "" && usearch == "" && irsearch == 0 {
		c.JSON(200, gin.H{
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": ci,
	})
}

// ListCarInspection handles request to get a list of carinspection entities
// @Summary List carinspection entities
// @Description list carinspection entities
// @ID list-carinspection
// @Produce json
// @Success 200 {array} ent.CarInspection
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carinspections [get]
func (ctl *CarInspectionController) ListCarInspection(c *gin.Context) {
	cis, err := ctl.client.CarInspection.
		Query().
		WithAmbulance().
		WithInspectionresult().
		WithUser().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cis)
}

// DeleteCarInspection handles DELETE requests to delete a carinspection entity
// @Summary Delete a carinspection entity by ID
// @Description get carinspection by ID
// @ID delete-carinspection
// @Produce  json
// @Param id path int true "CarInspection ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carinspections/{id} [delete]
func (ctl *CarInspectionController) DeleteCarInspection(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.CarInspection.
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

// NewCarInspectionController creates and registers handles for the carinspection controller
func NewCarInspectionController(router gin.IRouter, client *ent.Client) *CarInspectionController {
	ci := &CarInspectionController{
		client: client,
		router: router,
	}
	ci.register()
	return ci
}

// InitCarInspectionController registers routes to the main engine
func (ctl *CarInspectionController) register() {
	cis := ctl.router.Group("/carinspections")
	cisu := ctl.router.Group("/searchcarinspections")

	cis.GET("", ctl.ListCarInspection)

	// CRUD
	cis.POST("", ctl.CreateCarInspection)
	cis.GET(":id", ctl.GetCarInspection)
	cis.DELETE(":id", ctl.DeleteCarInspection)

	cisu.GET("", ctl.GetCarInspectionBySearch)
}
