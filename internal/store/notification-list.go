package store

import "github.com/johnearl92/xendit-notification-service.git/internal/model"

// NewNotificationList NotificationList provider
func NewNotificationList() *NotificationList {
	list := &NotificationList{}
	list.Init()
	return list
}

// NotificationList list definition
type NotificationList struct {
	BaseList
	items []*model.Notification
}

// Model model to use for db
func (p *NotificationList) Model() model.Model {
	return &model.Notification{}
}

// ItemsPtr pointer items in the list
func (p *NotificationList) ItemsPtr() interface{} {
	return &p.items
}

// Items items in the list
func (p *NotificationList) Items() []*model.Notification {
	return p.items
}
