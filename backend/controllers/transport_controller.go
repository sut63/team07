package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/transport"
	"github.com/team07/app/ent/user"
	"github.com/team07/app/ent/receive"
	"github.com/team07/app/ent/send"
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

	s, err := ctl.client.Send.
		Query().
		Where(send.IDEQ(int(obj.SendID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "send not found",
		})
		return
	}

	r, err := ctl.client.Receive.
		Query().
		Where(receive.IDEQ(int(obj.ReceiveID))).
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
		SetSendid(s).
		SetReceiveid(r).
		SetUser(u).
		SetAmbulance(ambulance).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ts)
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
		WithSendid().
		WithReceiveid().
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
		WithSendid().
		WithReceiveid().
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

	transports.GET("", ctl.ListTransport)

	// CRUD
	transports.GET(":id", ctl.GetTransport)
	//transport.POST("",ctl.CreateTransport)
	transports.DELETE(":id", ctl.DeleteTransport)
}

