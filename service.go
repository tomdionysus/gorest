package gorest

import (
	// "errors"
)

type Service interface {
	HasName
	HasLogger
	HasDatastore
	HasModel
	HasPresenter

	Create(data *[]byte) (error, *[]byte)
	Read(id interface{}) (error, *[]byte)
	Find(query map[string]interface{}) (error, *[]byte)
	Update(id interface{}, data *[]byte) (error, *[]byte)
	Delete(id interface{}) (error)
}