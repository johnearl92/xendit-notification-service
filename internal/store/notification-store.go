package store

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/johnearl92/xendit-notification-service.git/internal/model"
	log "github.com/sirupsen/logrus"
)

// NewNotificationStore ...
func NewNotificationStore(db *gorm.DB) NotificationStore {
	return &notificationStore{BaseStore{
		DB: db,
	}}
}

type notificationStore struct {
	BaseStore
}

// NotificationStore ...
type NotificationStore interface {
	Create(notification *model.Notification, opts GetOpts) error
	Update(notification *model.Notification) error
	Get(id string, opts GetOpts) (*model.Notification, error)
}

func (p *notificationStore) Create(notification *model.Notification, opts GetOpts) error {
	notification.MerchantKey = uuid.New().String()
	db := p.DB.Create(notification)
	err := db.Error

	return err
}

func (p *notificationStore) Update(notification *model.Notification) error {
	db := p.BaseStore.Update(notification)
	return db
}

func (p *notificationStore) Get(id string, opts GetOpts) (*model.Notification, error) {
	db := p.DB.Where("id = ?", id)
	acct, err := p.Find(db, &model.Notification{}, opts)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return acct.(*model.Notification), nil
}
