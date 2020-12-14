// Package service Official-Receipt API.
package service

import (
	"github.com/johnearl92/xendit-notification-service.git/internal/model"
	"github.com/johnearl92/xendit-notification-service.git/internal/store"
)

// xenditService EOR service definition
type xenditService struct {
	notificationStore store.NotificationStore
}

// NewXenditService xenditService provider
func NewXenditService(notificationStore store.NotificationStore) XenditService {
	return &xenditService{
		notificationStore: notificationStore,
	}
}

// XenditService Receiptservice interface
type XenditService interface {
	CreateNotification(notification *model.Notification, opts store.GetOpts) error
	UpdateNotification(notification *model.Notification) error
	GetNotification(id string, opts store.GetOpts) (*model.Notification, error)
	SendNotification(url string, payload string)
}

func (p *xenditService) SendNotification(url string, payload string) {


}

// Create saves receipts in database
func (p *xenditService) CreateNotification(notification *model.Notification, opts store.GetOpts) error {
	return p.notificationStore.Create(notification, opts)
}

// Update updates the specific receipt
func (p *xenditService) UpdateNotification(notification *model.Notification) error {
	return p.notificationStore.Update(notification)
}

// Get find a specific receipt in database
func (p *xenditService) GetNotification(id string, opts store.GetOpts) (*model.Notification, error) {
	return p.notificationStore.Get(id, opts)
}
