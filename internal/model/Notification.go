package model

type Notification struct {
	BaseModel
	Signature    string
	Payload      string
	PaymentType  int
	MerchantKey  string
	MerchantCode string
	RefNo        string
	IsProcessed  bool
	IsDeleted    bool
}

type NotificationRequest struct {
	Signature    string
	Payload      interface{}
	PaymentType  int
	MerchantKey  string
	MerchantCode string
	RefNo        string
}

type CardType struct {
	CardOwnerfirstName string
	CardOwnerlastName string
	Amount float64
}

// NotificationReq parameters model
//
// swagger:parameters NotificationReq
type NotificationReq struct {
	// in: body
	// required: true
	Notification NotificationRequest
}
