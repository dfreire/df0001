package commands

import (
	"github.com/dfreire/df0001/model"
	"github.com/jinzhu/gorm"
	"labix.org/v2/mgo/bson"
)

type SignupCustomerWithNewsletterRequestData struct {
	Name   string `json:"name,omitempty"`
	Email  string `json:"email"`
	RoleId string `json:"roleId"`
}

func SignupCustomerWithNewsletter(db *gorm.DB, reqData SignupCustomerWithNewsletterRequestData) error {
	toFind := model.Customer{
		Email: reqData.Email,
	}

	toCreate := model.Customer{
		ID:    bson.NewObjectId().Hex(),
		Email: reqData.Email,
	}

	err := db.Where(toFind).FirstOrCreate(&toCreate).Error
	if err != nil {
		return err
	}

	toUpdate := model.Customer{
		Name:                 reqData.Name,
		RoleId:               reqData.RoleId,
		SignedUpToNewsletter: true,
	}

	return db.Model(&toUpdate).Updates(toUpdate).Error
}
