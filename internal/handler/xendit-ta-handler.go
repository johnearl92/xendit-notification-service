// Package handler Official-Receipt API.
//
// Official-Receipt service API endpoints
//
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/johnearl92/xendit-notification-service.git/internal/model"
	"github.com/johnearl92/xendit-notification-service.git/internal/service"
	"github.com/johnearl92/xendit-notification-service.git/internal/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//XenditHandler main handler for the eor
type XenditHandler struct {
	xenditService service.XenditService
}

// NewORHandler provides XenditHandler definition
func NewXenditHandler(p service.XenditService) *XenditHandler {
	return &XenditHandler{
		xenditService: p,
	}
}

// GetHeartBeat check if the server is up
func (h *XenditHandler) GetHeartBeat(res http.ResponseWriter, req *http.Request) {
	log.Debugln("invoke getHeartBeat")
	resp := &model.GenericResponse{
		Success: true,
	}

	utils.WriteEntity(res, http.StatusOK, resp)
	log.Debugln("end getHeartBeat")
}

func (h *XenditHandler) AddNotification(res http.ResponseWriter, req *http.Request) {
	log.Debugln("invoke AddNotification")
	var data model.NotificationRequest
	log.Infoln("Reading Request")
	if err := utils.ReadEntity(req, &data); err != nil {
		log.Errorln(err)
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}

	//TODO get account details and verify signature

	var account = &model.Notification{
		MerchantKey:  data.MerchantKey,
		MerchantCode: data.MerchantCode,
		IsDeleted:    false,
		Payload:      data.Payload,
		PaymentType:  data.PaymentType,
		IsProcessed:  false,
		RefNo:        data.RefNo,
		Signature:    data.Signature,
	}

	if err := h.xenditService.CreateNotification(account, nil); err != nil {
		log.Error(err.Error())
		utils.WriteServerError(res, "/account", "Failed to save notification",
			fmt.Sprintf("Failed to save account. Please contact the administrator. Error: %s", err.Error()))
		return
	}

	resp := &model.GenericResponse{
		Success: true,
	}
	log.Debugln("end addComment")
	utils.WriteEntity(res, http.StatusOK, resp)

}

func (h *XenditHandler) TestNotification(res http.ResponseWriter, req *http.Request) {
	log.Debugln("invoke AddNotification")
	var data model.NotificationRequest
	log.Infoln("Reading Request")
	if err := utils.ReadEntity(req, &data); err != nil {
		log.Errorln(err)
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}



	resp := &model.GenericResponse{
		Success: true,
	}
	log.Debugln("end addComment")
	utils.WriteEntity(res, http.StatusOK, resp)

}

// Register registers the route
func (h *XenditHandler) Register(router *mux.Router) {

	// swagger:operation GET  /health GenericRes
	// ---
	// summary: This will check if the server is up
	// responses:
	//   "200":
	//     "$ref": "#/responses/GenericRes"
	//   "400":
	//     "$ref": "#/responses/JSONErrors"
	//   "500":
	//     "$ref": "#/responses/JSONErrors"
	router.HandleFunc("/health", h.GetHeartBeat).Methods(http.MethodGet)
	log.Info("[GET] /health registered")

	// swagger:operation POST /notification notif NotificationReq
	// Add notification
	// ---
	// responses:
	//   "200":
	//     "$ref": "#/responses/GenericRes"
	//   "400":
	//     "$ref": "#/responses/JSONErrors"
	//   "500":
	//     "$ref": "#/responses/JSONErrors"
	router.HandleFunc("/notification", h.AddNotification).Methods(http.MethodPost)
	log.Info("[POST] /notification registered")


	// swagger:operation POST /notification/test notif NotificationReq
	// Add notification
	// ---
	// responses:
	//   "200":
	//     "$ref": "#/responses/GenericRes"
	//   "400":
	//     "$ref": "#/responses/JSONErrors"
	//   "500":
	//     "$ref": "#/responses/JSONErrors"
	router.HandleFunc("/notification/test", h.TestNotification).Methods(http.MethodPost)
	log.Info("[POST] /notification/test registered")



}
