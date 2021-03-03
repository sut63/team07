package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/transport"
	"github.com/team07/app/ent/user"
	"github.com/team07/app/ent/hospital"
	"github.com/team07/app/ent/ambulance"
)

// TransportController defines the struct for the transport controller
type TransportController struct {
	client *ent.Client
	router gin.IRouter
}
type Transport struct {
	SendID int
	ReceiveID int
	UserID int
	AmbulanceID int
	Symptom string
	Drugallergy string
	Note string
}
// CreateTransport handles POST requests for adding transport entities
// @Summary Create transport
// @Description Create transport
// @ID create-transport
// @Accept   json
// @Produce  json
// @Param transport body Transport true "Transport entity"
// @Success 200 {object} ent.Transport
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /transports [post]
func (ctl *TransportController) CreateTransport(c *gin.Context) {
	obj := Transport{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "transport binding failed",
		})
		return
	}

	s, err := ctl.client.Hospital.
		Query().
		Where(hospital.IDEQ(int(obj.SendID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "send not found",
		})
		return
	}

	r, err := ctl.client.Hospital.
		Query().
		Where(hospital.IDEQ(int(obj.ReceiveID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "receive not found",
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
	ambulance, err := ctl.client.Ambulance.
		Query().
		Where(ambulance.IDEQ(int(obj.AmbulanceID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ambulance not found",
		})
		return
	}

	ts, err := ctl.client.Transport.
		Create().
		SetSend(s).
		SetReceive(r).
		SetDrugallergy(obj.Drugallergy).
		SetNote(obj.Note).
		SetSymptom(obj.Symptom).
		SetUser(u).
		SetAmbulance(ambulance).
		Save(context.Background())
	if err != nil {
		
		c.JSON(400, gin.H{
			"status": false,
			"error": err,

		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": ts,
	})
}

// GetTransport handles GET requests to retrieve a transport entity
// @Summary Get a transport entity by ID
// @Description get transport by ID
// @ID get-transport
// @Produce  json
// @Param id path int true "Transport ID"
// @Success 200 {object} ent.Transport
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /transports/{id} [get]
func (ctl *TransportController) GetTransport(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ts, err := ctl.client.Transport.
		Query().
		WithSend().
		WithReceive().
		WithAmbulance().
		WithUser().
		Where(transport.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ts)
}
// GetTransportBySearch handles GET requests to retrieve a transport entity
// @Summary Get a transport entity by Ambulance
// @Description get transport by Ambulance
// @ID get-transport-by-ambulance
// @Produce  json
// @Param ambulance query string false "Ambulance Search"
// @Param send query int false "Send Search"
// @Param receive query string false "Receive Search"
// @Success 200 {object} ent.Transport
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /searchtransports [get]
func (ctl *TransportController) GetTransportBySearch(c *gin.Context) {
	asearch := c.Query("ambulance")
	ssearch, err := strconv.ParseInt(c.Query("send"), 10, 64)
	rsearch, err := strconv.ParseInt(c.Query("receive"), 10, 64)

	sstring := ""
	s, err := ctl.client.Hospital.
		Query().
		Where(hospital.IDEQ(int(ssearch))).
		Only(context.Background())

	if s != nil {
		sstring = s.Hospital
	}

	rstring := ""
	r, err := ctl.client.Hospital.
		Query().
		Where(hospital.IDEQ(int(rsearch))).
		Only(context.Background())

	if r != nil {
		rstring = r.Hospital
	}

	ts, err := ctl.client.Transport.
		Query().
		WithAmbulance().
		WithSend().
		WithReceive().
		WithUser().
		Where(transport.HasSendWith(hospital.HospitalContains(sstring))).
		Where(transport.HasReceiveWith(hospital.HospitalContains(rstring))).
		Where(transport.HasAmbulanceWith(ambulance.CarregistrationContains(asearch))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	if asearch == "" && ssearch == 0 && rsearch == 0 {
		c.JSON(200, gin.H{
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"data":    ts,
	})
}
// ListTransport handles request to get a list of transport entities
// @Summary List transport entities
// @Description list transport entities
// @ID list-transport
// @Produce json
// @Success 200 {array} ent.Transport
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /transports [get]
func (ctl *TransportController) ListTransport(c *gin.Context) {
	transports, err := ctl.client.Transport.
		Query().
		WithSend().
		WithReceive().
		WithAmbulance().
		WithUser().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, transports)
}

// DeleteTransport handles DELETE requests to delete a transport entity
// @Summary Delete a transport entity by ID
// @Description get transport by ID
// @ID delete-transport
// @Produce  json
// @Param id path int true "Transport ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /transports/{id} [delete]
func (ctl *TransportController) DeleteTransport(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Transport.
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

// NewTransportController creates and registers handles for the transport controller
func NewTransportController(router gin.IRouter, client *ent.Client) *TransportController {
	transport := &TransportController{
		client: client,
		router: router,
	}
	transport.register()
	return transport
}

// InitTransportController registers routes to the main engine
func (ctl *TransportController) register() {
	transports := ctl.router.Group("/transports")
	searchtransport := ctl.router.Group("/searchtransports")

	transports.GET("", ctl.ListTransport)

	// CRUD
	transports.GET(":id", ctl.GetTransport)
	transports.POST("",ctl.CreateTransport)
	transports.DELETE(":id", ctl.DeleteTransport)
	
	searchtransport.GET("", ctl.GetTransportBySearch)
}

